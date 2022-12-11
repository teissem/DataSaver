package compression

import (
	"os"

	"teissem.fr/data_saver/src/configuration"
)

func Compress(configuration *configuration.Configuration) error {
	switch configuration.Compression {
	case "zip":
		CompressZip(configuration.Destination, configuration.Destination+"."+configuration.Compression)
		os.RemoveAll(configuration.Destination)
		return nil
	case "tar":
		CompressTar(configuration.Destination, configuration.Destination+"."+configuration.Compression)
		os.RemoveAll(configuration.Destination)
		return nil
	default:
		return nil
	}
}
