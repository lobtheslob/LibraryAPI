package main

import (
	"testing"
)

func TestCreateBook(t *testing.T) {


	tests := []struct{
		
		jsonStr []byte 	

	}{{
		jsonStr : []byte(`{"id":4,"name":"a name", "author":"some author","isbn":"1234567890-90897-12912021"}`),
    }}

    for _, test := range tests {
	expected := `{"id":4,"name":"a name", "author":"some author","isbn":"1234567890-90897-12912021"}`
	 
    }
}