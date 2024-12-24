package sqlites

import dklserrors "dkl.dklsa.mailer/iternal/dkls_errors"

type UrlTypeStorage struct {
	Tables
}

func NewTablesUrlType() UrlTypeStorage {
	return UrlTypeStorage{}
}

func (s UrlTypeStorage) New() error {
	return dklserrors.NotRelizedError()
}

func (s UrlTypeStorage) Get(id string) (Storage, error) {
	return Storage{}, dklserrors.NotRelizedError()
}

func (s UrlTypeStorage) Save(st Storage) error {
	return dklserrors.NotRelizedError()
}

func (s UrlTypeStorage) Delete(st Storage) error {
	return dklserrors.NotRelizedError()
}

func (s UrlTypeStorage) Update(st Storage) error {
	return dklserrors.NotRelizedError()
}
