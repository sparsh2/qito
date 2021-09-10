package encryption

import (
	"reflect"
	"testing"
)

func TestEncryptionXOR_encrypt(t *testing.T) {
	type fields struct {
		passphrase string
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "should encrypt",
			fields: fields{
				passphrase: "secret_password",
			},
			args: args{
				bytes: []byte("plaintext"),
			},
			want:    []byte{3, 9, 2, 27, 11, 0, 58, 8, 21},
			wantErr: false,
		},
		{
			name: "should throw error for empty passphrase",
			fields: fields{
				passphrase: "",
			},
			args: args{
				bytes: []byte("plaintext"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EncryptionServiceXOR{
				passphrase: tt.fields.passphrase,
			}
			got, err := e.encrypt(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptionXOR.encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncryptionXOR.encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptionServiceXOR_decrypt(t *testing.T) {
	type fields struct {
		passphrase string
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "should decrypt",
			fields: fields{
				passphrase: "secret_password",
			},
			args: args{
				bytes: []byte{3, 9, 2, 27, 11, 0, 58, 8, 21, 73, 83, 4, 0, 31, 1, 7, 13, 10, 28, 2, 84, 44, 21, 2, 1, 22, 3, 79, 6, 11, 83, 0, 13, 17, 23, 13, 47, 4},
			},
			want:    []byte("plaintext: something secret to encrypt"),
			wantErr: false,
		},
		{
			name: "should throw error for empty passphrase",
			fields: fields{
				passphrase: "",
			},
			args: args{
				bytes: []byte{3, 9, 2, 27, 11, 0, 58, 8, 21},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EncryptionServiceXOR{
				passphrase: tt.fields.passphrase,
			}
			got, err := e.decrypt(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptionServiceXOR.decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncryptionServiceXOR.decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
