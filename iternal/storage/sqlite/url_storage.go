package sqlites

import dklserrors "dkl.dklsa.mailer/iternal/dkls_errors"

type UrlStorages struct {
	Tables
}

func NewTablesUrl() UrlStorages {
	return UrlStorages{}
}

func (s UrlStorages) New() error {
	return dklserrors.NotRelizedError()
}

func (s UrlStorages) Get(id string) (Storage, error) {
	return Storage{}, dklserrors.NotRelizedError()
}

func (s UrlStorages) Save(st Storage) error {
	return dklserrors.NotRelizedError()
}

func (s UrlStorages) Delete(st Storage) error {
	return dklserrors.NotRelizedError()
}

func (s UrlStorages) Update(st Storage) error {
	return dklserrors.NotRelizedError()
}
