package main

import (
  "testing"
)

func TestBasic(t *testing.T) {

// Basic CRUS Tests for User
  t.Run("User", func(t *testing.T){

    var TestUserId = "orcid:1234-1234-1234-1234"
    var TestUser User


    TestUser.Id = "orcid:1234-1234-1234-1234"
    TestUser.Delete()

    TestUser = User{
      Id: "orcid:1234-1234-1234-1234",
      Name: "Joe Schmoe",
      Email: "JoeSchmoe@example.org",
      IsAdmin: false,
      Groups: []string{},
    }


    t.Run("Create", func(t *testing.T){

      err := TestUser.Create()

      if err != nil {
        t.Fatalf("Failed to Create the User: %s", err.Error())
      }

    })

    t.Run("Get", func(t *testing.T){

      findUser := User{Id: TestUserId}
      err := findUser.Get()


      if err != nil {
        t.Fatalf("Failed to Get User: %s", err.Error())
      }

      t.Logf("Found User: %+v", findUser)

    })

    t.Run("List", func(t *testing.T){
      userList, err := listUsers()

      if err != nil {
        t.Fatalf("Failed to List Users: %s", err.Error())
      }

      if len(userList) == 0 {
        t.Fatalf("Failed to List any Users")
      }

      t.Logf("Found Users: %+v", userList)

    })

    t.Run("Delete", func(t *testing.T){

      delUser := User{Id: TestUserId}
      err := delUser.Delete()

      if err != nil {
        t.Fatalf("Failed to Delete User: %s", err.Error())
      }

      t.Logf("Deleted User: %+v", delUser)

    })

  })

  // Basic CRUD Tests for Groups
  t.Run("Group", func(t *testing.T) {

      admin := User{Id: "orcid:1", Groups: []string{}}
      admin.Create()


      user := User{Id: "orcid:2", Groups: []string{}}
      user.Create()

      g := Group{Id: "group1", Admin:"orcid:1", Members: []string{"orcid:2"}}


    t.Run("Create", func(t *testing.T){

      var err error

      if err = g.Create(); err != nil {
        t.Fatalf("Failed to Create Group: %s", err.Error())
      }

      // Verify Admin has member of group
      var updatedAdmin User
      updatedAdmin.Id = admin.Id
      err = updatedAdmin.Get()

      if err != nil {
        t.Fatalf("Failed to Fetch Updated Admin: %s", err.Error())
      }

      if len(updatedAdmin.Groups) != 1 {
        t.Fatalf("Admin is not listed as member of group: %+v", updatedAdmin)
      }


      // Verify User is member of group
      var updatedUser User
      updatedUser.Id = user.Id
      err = updatedUser.Get()

      if err != nil {
        t.Fatalf("Failed to Fetch Updated User: %s", err.Error())
      }

      if len(updatedUser.Groups) != 1 {
        t.Fatalf("Admin is not listed as member of group: %+v", updatedUser)
      }

    })

    t.Run("Get", func(t *testing.T){
      found := Group{Id: "group1"}
      err := found.Get()
      if err != nil {
        t.Fatalf("Failed to Find Group: %s", err.Error())
      }
    })

    t.Run("List", func(t *testing.T){
      _, err := listGroups()
      if err != nil {
        t.Fatalf("Failed to List Groups: %s", err.Error())
      }
    })

    t.Run("Delete", func(t *testing.T){
      del := Group{Id: "group1"}
      err := del.Delete()
      if err != nil {
        t.Fatalf("Failed to Delete Group: %s", err.Error())
      }

      t.Logf("Deleted Group: %+v", del)
    })


  })


  // Basic CRUD Tests for Resources
  t.Run("Resource", func(t *testing.T) {

    u := User{Id: "owner"}

    r := Resource{Id: "res1", Owner: u.Id}

    t.Run("Create", func(t *testing.T){
      err := r.Create()
      if err != nil {
        t.Fatalf("Failed to Create Resource: %s", err.Error())
      }
    })

    t.Run("Get", func(t *testing.T){
      found := Resource{Id: "res1"}
      err := found.Get()
      if err != nil {
        t.Fatalf("Failed to Find Resource: %s", err.Error())
      }
    })

    t.Run("List", func(t *testing.T){
      rlist, err := listResources()
      if err != nil {
        t.Fatalf("Failed to List Resources: %s", err.Error())
      }
      t.Logf("ListResources: %+v", rlist)
    })

    t.Run("Delete", func(t *testing.T){
      del := Resource{Id: "res1"}
      err := del.Delete()
      if err != nil {
        t.Fatalf("DeleteResourceError: %s", err.Error())
      }
      t.Logf("DeleteResource: %+v", del)
    })

  })


  // Basic CRUD Tests Policy
  t.Run("Policy", func(t *testing.T) {

    owner := User{Id: "owner"}
    owner.Create()

    u := User{Id: "u1"}
    u.Create()

    r := Resource{Id: "r1", Owner: "owner"}
    r.Create()

    p := Policy{Id: "p1", Resource: "r1"}

    t.Run("Create", func(t *testing.T){
      err := p.Create()
      if err != nil {
        t.Fatalf("CreatePolicyErr: %s", err.Error())
      }
    })

    t.Run("Get", func(t *testing.T){
      found := Policy{Id:"p1"}
      err := found.Get()
      if err != nil {
        t.Fatalf("ERROR FindPolicy: %s", err.Error())
      }

      t.Logf("INFO FindPolicy: %+v", found)
    })

    t.Run("List", func(t *testing.T){
      plist, err := listPolicies()
      if err != nil {
        t.Fatalf("ERROR ListPolicy: %s", err.Error())
      }
      t.Logf("INFO ListPolicy: %+v", plist)
    })

    t.Run("Delete", func(t *testing.T){
      del := Policy{Id: "p1"}
      err := del.Delete()
      if err != nil {
        t.Fatalf("ERROR DeletePolicy: %s", err.Error())
      }

      t.Logf("INFO DeletePolicy: +%v", del)
    })

    owner.Delete()
    u.Delete()
    r.Delete()

  })


  // Basic CRUD Tests for Challneges
  t.Run("Challenge", func(t *testing.T) {

    owner := User{Id: "owner"}
    owner.Create()

    u := User{Id: "u1"}
    u.Create()

    r := Resource{Id: "r1", Owner: "owner"}
    r.Create()

    p := Policy{
      Id: "p1",
      Resource: "r1",
      Principal: []string{"u1"},
      Action: []string{"read"},
      Effect: "Allow",
    }
    p.Create()



    t.Run("Evaluate", func(t *testing.T){

      t.Run("Owner", func(t *testing.T){

        c := Challenge{
          Principal: "owner",
          Resource: "r1",
          Action: "delete",
        }

        c.Evaluate()

        if !c.Granted {
          t.Fatalf("ERROR ChallengeOwner: Owner of Resource Wrongly Denied Permission")
        }
      })

      t.Run("PolicyAllowed", func(t *testing.T){
        c := Challenge{
          Principal: "u1",
          Resource: "r1",
          Action: "read",
        }

        err := c.Evaluate()
        if err != nil {
          t.Fatalf("ERROR ChallengeEvaluation: %s", err.Error())
        }

        if !c.Granted {
          t.Fatalf("ERROR Challenge Incorrectly Denied")
        }

      })

      t.Run("PolicyMissingAction", func(t *testing.T){

        c := Challenge{
          Principal: "u1",
          Resource: "r1",
          Action: "write",
        }

        err := c.Evaluate()
        if err != nil {
          t.Fatalf("ERROR ChallengeEvaluation: %s", err.Error())
        }

        if c.Granted {
          t.Fatalf("ERROR Challenge Incorrectly Granted")
        }

      })

    })

    t.Run("List", func(t *testing.T){
      clist, err := listChallenges()
      if err != nil {
        t.Fatalf("ERROR ListChallenges: %s", err.Error())
      }
      t.Logf("INFO ListChallenges: %+v", clist)

    })


  })

}