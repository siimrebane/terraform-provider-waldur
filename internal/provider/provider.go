package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/waldur/terraform-waldur-provider/internal/resources"
	"github.com/waldur/terraform-waldur-provider/internal/datasources"
	"github.com/waldur/terraform-waldur-provider/internal/client"
)

// Ensure waldurProvider satisfies various provider interfaces.
var _ provider.Provider = &waldurProvider{}

// waldurProvider defines the provider implementation.
type waldurProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance testing.
	version string
}

// waldurProviderModel describes the provider data model.
type waldurProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Token    types.String `tfsdk:"token"`
}

func (p *waldurProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "waldur"
	resp.Version = p.version
}

func (p *waldurProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Waldur API endpoint URL",
				Required:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "API authentication token",
				Required:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *waldurProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data waldurProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API client
	apiClient, err := client.NewClient(&client.Config{
		Endpoint: data.Endpoint.ValueString(),
		Token:    data.Token.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create API Client",
			"An error occurred when creating the API client: "+err.Error(),
		)
		return
	}

	// Make client available to resources and data sources
	resp.DataSourceData = apiClient
	resp.ResourceData = apiClient
}

func (p *waldurProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewStructureProjectResource,
		resources.NewOpenstackInstanceResource,
	}
}

func (p *waldurProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewStructureProjectDataSource,
		datasources.NewOpenstackFlavorDataSource,
		datasources.NewOpenstackImageDataSource,
		datasources.NewOpenstackTenantDataSource,
		datasources.NewOpenstackInstanceDataSource,
		datasources.NewOpenstackSecurityGroupDataSource,
		datasources.NewOpenstackFloatingIpDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &waldurProvider{
			version: version,
		}
	}
}
