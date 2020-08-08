package cmd

import (
	"fmt"
	"jumpadmin/internal/service"
	"strings"

	"github.com/CodyGuo/glog"
	"github.com/spf13/cobra"
)

var (
	organizations []string
	users         []string
	admins        []string
	list          bool
)

var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "organization manager",
	Long:  "add organization user admin",
	Run: func(cmd *cobra.Command, args []string) {
		if list {
			svc := service.New(ctx)
			organizations, err := svc.ListOrganization()
			if err != nil {
				glog.Fatal(err)
			}
			for _, o := range organizations {
				fmt.Println(o.ID, o.Name)
			}
		} else {
			cmd.Help()
		}
	},
}

var orgAddCmd = &cobra.Command{
	Use:   "add",
	Short: "org add -o org1,org2",
	Long:  "org add -o org1,org2 [-u user1,user2] [-a admin1, admin2]",
	Run: func(cmd *cobra.Command, args []string) {
		svc := service.New(ctx)
		for _, organization := range organizations {
			if len(admins) == 0 && len(users) == 0 {
				err := svc.CreateOrganization(&service.CreateOrganizationRequest{Name: organization})
				if err != nil {
					glog.Errorf("add org failed, org: %s, error: %v", organization, err)
				} else {
					glog.Infof("add org success, org: %s", organization)
				}
			}
			for _, u := range admins {
				err := svc.CreateOrganizationAdmins(&service.OrganizationAdminsRequest{Organization: organization, Username: u})
				if err != nil {
					if !strings.Contains(err.Error(), "Duplicate entry") {
						glog.Errorf("add admin failed, org: %s, user: %s, error: %v", organization, u, err)
					} else {
						glog.Errorf("admin %s already exists in the %s org", u, organization)
					}
				} else {
					glog.Infof("add admin success, org: %s user: %s", organization, u)
				}
			}
			for _, u := range users {
				err := svc.CreateOrganizationUsers(&service.OrganizationUsersRequest{Organization: organization, Username: u})
				if err != nil {
					if !strings.Contains(err.Error(), "Duplicate entry") {
						glog.Errorf("add user failed, org: %s, user: %s, error: %v", organization, u, err)
					} else {
						glog.Errorf("user %s already exists in the %s org", u, organization)
					}
				} else {
					glog.Infof("add user success, org: %s user: %s", organization, u)
				}
			}
		}
	},
}

var orgDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "org delete -o org1,org2 [-u user1,user2] [-a admin1, admin2]",
	Long:  "org delete -o org1,org2 [-u user1,user2] [-a admin1, admin2]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(users) == 0 && len(admins) == 0 {
			cmd.Help()
			return
		}
		svc := service.New(ctx)
		for _, organization := range organizations {
			for _, u := range admins {
				err := svc.DeleteOrganizationAdmins(&service.OrganizationAdminsRequest{Organization: organization, Username: u})
				if err != nil {
					glog.Errorf("delete admin failed, org: %s, user: %s, error: %v", organization, u, err)
				} else {
					glog.Infof("delete admin success, org: %s user: %s", organization, u)
				}
			}
			for _, u := range users {
				err := svc.DeleteOrganizationUsers(&service.OrganizationUsersRequest{Organization: organization, Username: u})
				if err != nil {
					glog.Errorf("delete user failed, org: %s, user: %s, error: %v", organization, u, err)
				} else {
					glog.Infof("delete user success, org: %s user: %s", organization, u)

				}
			}
		}
	},
}

func init() {
	orgCmd.AddCommand(orgAddCmd, orgDeleteCmd)
	orgCmd.Flags().BoolVarP(&list, "list", "l", false, "list organization")
	orgAddCmd.Flags().StringSliceVarP(&organizations, "orgs", "o", nil, "set organization")
	orgAddCmd.MarkFlagRequired("orgs")
	orgDeleteCmd.Flags().StringSliceVarP(&organizations, "orgs", "o", nil, "set organization")
	orgDeleteCmd.MarkFlagRequired("orgs")

	orgCmd.PersistentFlags().StringSliceVarP(&users, "user", "u", nil, "set user")
	orgCmd.PersistentFlags().StringSliceVarP(&admins, "admin", "a", nil, "set admin")
}
