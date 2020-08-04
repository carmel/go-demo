package test

import (
	"fmt"
	"testing"

	"github.com/mikespook/gorbac"
)

func TestAuth(t *testing.T) {
	//Get a new instance of RBAC
	rbac := gorbac.New()
	//Get some new roles
	rA := gorbac.NewStdRole("role-a")
	rB := gorbac.NewStdRole("role-b")
	rC := gorbac.NewStdRole("role-c")
	rD := gorbac.NewStdRole("role-d")
	rE := gorbac.NewStdRole("role-e")
	//Get some new permissions
	pA := gorbac.NewStdPermission("permission-a")
	pB := gorbac.NewStdPermission("permission-b")
	pC := gorbac.NewStdPermission("permission-c")
	pD := gorbac.NewStdPermission("permission-d")
	pE := gorbac.NewStdPermission("permission-e")
	//Add the permissions to roles
	rA.Assign(pA)
	rB.Assign(pB)
	rC.Assign(pC)
	rD.Assign(pD)
	rE.Assign(pE)
	//After initailization, add the roles to the RBAC instance
	rbac.Add(rA)
	rbac.Add(rB)
	rbac.Add(rC)
	rbac.Add(rD)
	rbac.Add(rE)
	//And set the inheritance
	rbac.SetParent("role-a", "role-b")
	rbac.SetParents("role-b", []string{"role-c", "role-d"})
	rbac.SetParent("role-e", "role-d")
	//Checking the permission
	if rbac.IsGranted("role-a", pA, nil) &&
		rbac.IsGranted("role-a", pB, nil) &&
		rbac.IsGranted("role-a", pC, nil) &&
		rbac.IsGranted("role-a", pD, nil) {
		fmt.Println("The role-a has been granted permis-a, b, c and d.")
	}
	//And there are some built-in util-functions: InherCircle, AnyGranted, AllGranted
	rbac.SetParent("role-c", "role-a")
	if err := gorbac.InherCircle(rbac); err != nil {
		fmt.Println("A circle inheratance ocurred.")
	}
}
