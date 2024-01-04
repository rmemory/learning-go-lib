package programming


type Interface interface {
	NewUuid(withoutHyphen bool) string
	DebugJWT(tokenString string) (string, string, error)
}

/* The following interface makes it easier to mock methods that implement any of the interfaces in this file. We do this by using the following structure as a receiver type for any method implementing of the interfaces. That allows any testing code to have two different implementations: A real one and a mock.

(step 1) Attach this type as a receiver to any function that implements an interface in this file. For example, 

func (pf *ProgrammingFunctions) NewUuid(withoutHyphen bool) string {

(step 2) Modify test code by declaring the following variable, probably at the file scope outside of any functions

var pf ProgrammingFunctions = ProgrammingFunctions{}

(step 3) When calling the function (in this example, NewUuid) add the "pf" like follows:

uuidWithHyphen := pf.NewUuid(false)

(step 4) Clients of the interfaces (like learning-go-api, which uses the NewUuid function) will need to have the "p programming.Interface" passed to it so that it can use NewUuid. For example, 

func postUuid(p programming.Interface) {
	uuid := p.NewUuid(withoutHyphens)

Or, if clients are wanting to use mocked data (in client test code), they'd do something like this

func TestPostUuid(mockInterface *programminglib.MockInterface, t *testing.T) {
	mockInterface := programminglib.MockInterface{}
	mockCall := mockInterface.On("NewUuid", false)
	mockCall.Return("1ce44be5-fe68-46f7-a153-51c1c91a4ae4")
*/
type ProgrammingFunctions struct {
}