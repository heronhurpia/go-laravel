package celeritas

const version = "1.0.0"

type Celeritas struct {
	AppName string
	Debug   bool
	Version string
}

func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: {"handlers","migrations","views","data","public","tmp","logs","middleware"}
	}

	err := c.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil
}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// Cria pasta se não existir
		err := c.CreateDirIfNotExists((root + "/" + path))
		if err != nil {
			return err
		}
	}

	return nil
}
