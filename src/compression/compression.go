package compression

import (
	"os"

	"teissem.fr/data_saver/src/configuration"
)

func Compress(configuration *configuration.Configuration) error {
	switch configuration.Compression {
	case "zip":
		err := CompressZip(configuration.Destination, configuration.Destination+"."+configuration.Compression)
		if err != nil {
			return err
		}
		return os.RemoveAll(configuration.Destination)
	case "tar":
		err := CompressTar(configuration.Destination, configuration.Destination+"."+configuration.Compression)
		if err != nil {
			return err
		}
		return os.RemoveAll(configuration.Destination)
	default:
		return nil
	}
}
