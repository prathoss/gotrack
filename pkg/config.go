package pkg

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// generateCfgTemplate returns string template of configuration file
func generateCfgTemplate(cfg interface{}) (string, error) {
	cfgMap := make(map[string]interface{})
	if err := mapstructure.Decode(cfg, &cfgMap); err != nil {
		return "", err
	}
	out := &strings.Builder{}
	encoder := yaml.NewEncoder(out)
	if err := encoder.Encode(cfgMap); err != nil {
		return "", err
	}
	return out.String(), nil
}

// bindEnvVars scans structure for env tag and binds it with env variable
// parameter should be pointer
func bindEnvVars(cfg interface{}) error {
	tp := reflect.TypeOf(cfg)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	vtp := reflect.ValueOf(cfg)
	if vtp.Kind() == reflect.Ptr {
		vtp = vtp.Elem()
	}
	for i := 0; i < tp.NumField(); i++ {
		fld := tp.Field(i)
		vfld := vtp.Field(i)
		if vfld.Kind() == reflect.Ptr {
			vfld = vfld.Elem()
		}
		// it is struct, go scan fields and continue
		if vfld.Kind() == reflect.Struct {
			if err := bindEnvVars(vfld.Addr().Interface()); err != nil {
				return err
			}
			continue
		}

		envName := fld.Tag.Get("env")
		if envName == "" {
			continue
		}
		if !vfld.CanSet() {
			continue
		}
		if !viper.IsSet(envName) {
			continue
		}
		switch vfld.Kind() {
		case reflect.String:
			vfld.SetString(viper.GetString(envName))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			vfld.SetInt(viper.GetInt64(envName))
		case reflect.Bool:
			vfld.SetBool(viper.GetBool(envName))
		case reflect.Float64, reflect.Float32:
			vfld.SetFloat(viper.GetFloat64(envName))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			vfld.SetUint(viper.GetUint64(envName))
		}
	}
	return nil
}

// loadCfg binds config struct with viper, and validates it
// parameter should be pointer
func loadCfg(cfg interface{}) error {
	var errAgreg ErrorConfigErrors
	validate := validator.New()

	// load from viper cfg file
	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}
	if err := bindEnvVars(cfg); err != nil {
		return err
	}
	if err := validate.Struct(cfg); err != nil {
		errAgreg.AddValidationErrors(err, cfg)
	}
	if errAgreg.Any() {
		return errAgreg
	}
	return nil
}

type ErrorConfigErrors struct {
	msgs []string
}

func (e ErrorConfigErrors) Error() string {
	return fmt.Sprintf("Configuration errors:\n%s", strings.Join(e.msgs, "\n"))
}

func (e ErrorConfigErrors) Any() bool {
	return len(e.msgs) > 0
}

func (e *ErrorConfigErrors) AddValidationErrors(vErrs error, strct interface{}) {
	var errs validator.ValidationErrors
	tp := reflect.TypeOf(strct)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	if errors.As(vErrs, &errs) {
		for _, err := range errs {
			ctp := tp
			path := strings.Split(err.StructNamespace(), ".")
			var cfgFlPath []string
			var fld reflect.StructField
			// the first is name of current struct => not needed
			for i := 1; i < len(path); i++ {
				var ok bool
				fld, ok = ctp.FieldByName(path[i])
				if !ok {
					panic(fmt.Sprintf("could not find field %s in type %s", path[i], ctp.Name()))
				}
				ctp = fld.Type
				pathStep := fld.Tag.Get("mapstructure")
				if pathStep == "" {
					continue
				}
				cfgFlPath = append(cfgFlPath, pathStep)
			}
			cfgEnvNm := fld.Tag.Get("env")
			e.msgs = append(e.msgs, fmt.Sprintf("File config in path: '%s' or env var: '%s' is %s", strings.Join(cfgFlPath, "."), cfgEnvNm, err.Tag()))
		}
	}
}
