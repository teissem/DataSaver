# Architecture

The main architecture is composed of 3 modules

- configuration : Permits to parse configuration file
- datasource : Permits to get the data from different sources
- compression : Permits to compress the final product

## Configuration

The configuration package permits to parse the configuration file. For the moment, the only format available for the configuration file is JSON format.

To declare a new configuration format, you just have to create a new file in configuration package which create the struct configuration. The format of the function must be : `func(path string) (*Configuration, error)`.

After that, declare it in `SupportedConfigurationFormat` with the file extension waited and the function created.

## Datasource

The datasource package permits to get data from different sources. The currently available sources are :

- Filesystem (Folder)
- Git

## Compression

The compression package permits to compress the final result of the data saver. This permits to save a lot of space. The currently available compression algorithm are :

- zip
- tar
