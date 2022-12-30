// Package try simplifies the error handling in waiting Go 2 error handling.
//
// It is not the perfect solution, and it is not recommended to use in the libraries, but it can
// help for the simple apps.
//
// In Go it is common to handle errors like this:
//
//	func foo() error {
//	  v, err := bar()
//	  if err != nil {
//	    return err
//	  }
//
//	  if err := bar2(v); err != nil {
//	    return err
//	  }
//
//	  return nil
//	}
//
// It gets tedious when there are a lot of these checks. With the try package you can replace this
// code to:
//
//	func foo() (outErr error) {
//	  defer try.HandleAs(*outErr)
//
//	  v := try.ItVal(bar())
//
//	  try.It(bar2(v))
//
//	  return nil
//	}
package try
