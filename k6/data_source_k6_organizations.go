package k6

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOrganizations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOrganizationsRead,
		Schema: map[string]*schema.Schema{
			"organizations": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"logo": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"billing_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"billing_country": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"billing_email": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vat_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"created": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_saml_org": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

type Org struct {
	id              string
	name            string
	logo            string
	owner_id        string
	description     string
	billing_address string
	billing_country string
	billing_email   string
	vat_number      string
	created         string
	updated         string
	is_default      bool
	is_saml_org     bool
}

type OrgObj struct {
	Organizations []*Org
}

func dataSourceOrganizationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	token := m.(Config).Token

	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/organizations", "https://api.loadimpact.com/v3"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	var tokenString strings.Builder
	fmt.Fprintf(&tokenString, "token %s", token)
	req.Header.Set("Authorization", tokenString.String())

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	var o OrgObj
	err = json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("organizations", o.Organizations); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
