package sqlites

import dklserrors "dkl.dklsa.mailer/iternal/dkls_errors"

type DomainStorage struct {
	Tables
}

func NewTablesDomain() DomainStorage {
	return DomainStorage{}
}

func (s DomainStorage) New() error {
	return dklserrors.NotRelizedError()
}

func (s DomainStorage) Get(id string) (Storage, error) {
	return Storage{}, dklserrors.NotRelizedError()
}

func (s DomainStorage) Save(st Storage) error {
	return dklserrors.NotRelizedError()
}

func (s DomainStorage) Delete(st Storage) error {
	return dklserrors.NotRelizedError()
}

func (s DomainStorage) Update(st Storage) error {
	return dklserrors.NotRelizedError()
}
