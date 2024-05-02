// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"net/http"

	"github.com/go-faster/jx"

	ht "github.com/ogen-go/ogen/http"
)

func encodeCreateDepartmentRequest(
	req *CreateDepartmentReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateDeptEmployeeRequest(
	req *CreateDeptEmployeeReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateDeptManagerRequest(
	req *CreateDeptManagerReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateEmployeeRequest(
	req *CreateEmployeeReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateSalaryRequest(
	req *CreateSalaryReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateTitleRequest(
	req *CreateTitleReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeUpdateDepartmentRequest(
	req *UpdateDepartmentReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeUpdateDeptEmployeeRequest(
	req *UpdateDeptEmployeeReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeUpdateDeptManagerRequest(
	req *UpdateDeptManagerReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeUpdateEmployeeRequest(
	req *UpdateEmployeeReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeUpdateSalaryRequest(
	req *UpdateSalaryReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeUpdateTitleRequest(
	req *UpdateTitleReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}