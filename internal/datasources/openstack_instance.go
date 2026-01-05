package datasources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-waldur-provider/internal/client"
)

var _ = strings.Split // ensure strings is used

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OpenstackInstanceDataSource{}

func NewOpenstackInstanceDataSource() datasource.DataSource {
	return &OpenstackInstanceDataSource{}
}

// OpenstackInstanceDataSource defines the data source implementation.
type OpenstackInstanceDataSource struct {
	client *client.Client
}

// OpenstackInstanceDataSourceModel describes the data source data model.
type OpenstackInstanceDataSourceModel struct {
	UUID types.String `tfsdk:"id"`
	AttachVolumeUuid types.String `tfsdk:"attach_volume_uuid"`
	AvailabilityZoneName types.String `tfsdk:"availability_zone_name"`
	BackendId types.String `tfsdk:"backend_id"`
	CanManage types.Bool `tfsdk:"can_manage"`
	Customer types.String `tfsdk:"customer"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerName types.String `tfsdk:"customer_name"`
	CustomerNativeName types.String `tfsdk:"customer_native_name"`
	CustomerUuid types.String `tfsdk:"customer_uuid"`
	Description types.String `tfsdk:"description"`
	ExternalIp types.String `tfsdk:"external_ip"`
	Field types.String `tfsdk:"field"`
	Name types.String `tfsdk:"name"`
	NameExact types.String `tfsdk:"name_exact"`
	Page types.Int64 `tfsdk:"page"`
	PageSize types.Int64 `tfsdk:"page_size"`
	Project types.String `tfsdk:"project"`
	ProjectName types.String `tfsdk:"project_name"`
	ProjectUuid types.String `tfsdk:"project_uuid"`
	Query types.String `tfsdk:"query"`
	RuntimeState types.String `tfsdk:"runtime_state"`
	ServiceSettingsName types.String `tfsdk:"service_settings_name"`
	ServiceSettingsUuid types.String `tfsdk:"service_settings_uuid"`
	State types.String `tfsdk:"state"`
	Tenant types.String `tfsdk:"tenant"`
	TenantUuid types.String `tfsdk:"tenant_uuid"`
	Uuid types.String `tfsdk:"uuid"`
	AccessUrl types.String `tfsdk:"access_url"`
	Action types.String `tfsdk:"action"`
	AvailabilityZone types.String `tfsdk:"availability_zone"`
	ConnectDirectlyToExternalNetwork types.Bool `tfsdk:"connect_directly_to_external_network"`
	Cores types.Int64 `tfsdk:"cores"`
	Created types.String `tfsdk:"created"`
	Disk types.Int64 `tfsdk:"disk"`
	ErrorMessage types.String `tfsdk:"error_message"`
	ErrorTraceback types.String `tfsdk:"error_traceback"`
	ExternalAddress types.List `tfsdk:"external_address"`
	ExternalIps types.List `tfsdk:"external_ips"`
	FlavorDisk types.Int64 `tfsdk:"flavor_disk"`
	FlavorName types.String `tfsdk:"flavor_name"`
	FloatingIps types.List `tfsdk:"floating_ips"`
	HypervisorHostname types.String `tfsdk:"hypervisor_hostname"`
	ImageName types.String `tfsdk:"image_name"`
	InternalIps types.List `tfsdk:"internal_ips"`
	IsLimitBased types.Bool `tfsdk:"is_limit_based"`
	IsUsageBased types.Bool `tfsdk:"is_usage_based"`
	KeyFingerprint types.String `tfsdk:"key_fingerprint"`
	KeyName types.String `tfsdk:"key_name"`
	Latitude types.Float64 `tfsdk:"latitude"`
	Longitude types.Float64 `tfsdk:"longitude"`
	MarketplaceCategoryName types.String `tfsdk:"marketplace_category_name"`
	MarketplaceCategoryUuid types.String `tfsdk:"marketplace_category_uuid"`
	MarketplaceOfferingName types.String `tfsdk:"marketplace_offering_name"`
	MarketplaceOfferingUuid types.String `tfsdk:"marketplace_offering_uuid"`
	MarketplacePlanUuid types.String `tfsdk:"marketplace_plan_uuid"`
	MarketplaceResourceState types.String `tfsdk:"marketplace_resource_state"`
	MarketplaceResourceUuid types.String `tfsdk:"marketplace_resource_uuid"`
	MinDisk types.Int64 `tfsdk:"min_disk"`
	MinRam types.Int64 `tfsdk:"min_ram"`
	Modified types.String `tfsdk:"modified"`
	Ports types.List `tfsdk:"ports"`
	Ram types.Int64 `tfsdk:"ram"`
	ResourceType types.String `tfsdk:"resource_type"`
	SecurityGroups types.List `tfsdk:"security_groups"`
	ServerGroup types.Object `tfsdk:"server_group"`
	ServiceName types.String `tfsdk:"service_name"`
	ServiceSettings types.String `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String `tfsdk:"service_settings_error_message"`
	ServiceSettingsState types.String `tfsdk:"service_settings_state"`
	StartTime types.String `tfsdk:"start_time"`
	Url types.String `tfsdk:"url"`
	UserData types.String `tfsdk:"user_data"`
	Volumes types.List `tfsdk:"volumes"`
}

func (d *OpenstackInstanceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_instance"
}

func (d *OpenstackInstanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackInstance data source - lookup by name or UUID",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Resource UUID",
			},
			"attach_volume_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Filter for attachment to volume UUID",
			},
			"availability_zone_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Availability zone name",
			},
			"backend_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Backend ID",
			},
			"can_manage": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Can manage",
			},
			"customer": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer UUID",
			},
			"customer_abbreviation": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer abbreviation",
			},
			"customer_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer name",
			},
			"customer_native_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer native name",
			},
			"customer_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Customer UUID",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Description",
			},
			"external_ip": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "External IP",
			},
			"field": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name",
			},
			"name_exact": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Name (exact)",
			},
			"page": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "A page number within the paginated result set.",
			},
			"page_size": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Number of results to return per page.",
			},
			"project": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project UUID",
			},
			"project_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project name",
			},
			"project_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project UUID",
			},
			"query": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Search by name, internal IP, or external IP",
			},
			"runtime_state": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_settings_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Service settings name",
			},
			"service_settings_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Service settings UUID",
			},
			"state": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "State",
			},
			"tenant": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Tenant URL",
			},
			"tenant_uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Tenant UUID",
			},
			"uuid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "UUID",
			},
			"access_url":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"action":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"availability_zone":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Availability zone where this instance is located",
			},
			"connect_directly_to_external_network":schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "If True, instance will be connected directly to external network",
			},
			"cores":schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of cores in a VM",
			},
			"created":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"disk":schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Disk size in MiB",
			},
			"error_message":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"error_traceback":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"external_address":schema.ListAttribute{
				ElementType: types.StringType,
				Computed:            true,
				MarkdownDescription: "",
			},
			"external_ips":schema.ListAttribute{
				ElementType: types.StringType,
				Computed:            true,
				MarkdownDescription: "",
			},
			"flavor_disk":schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Flavor disk size in MiB",
			},
			"flavor_name":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the flavor used by this instance",
			},
			"floating_ips":schema.ListAttribute{
				// Flatten complex objects to list of strings for simpler schema
				ElementType: types.StringType,
				Computed:            true,
				MarkdownDescription: "",
			},
			"hypervisor_hostname":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the hypervisor hosting this instance",
			},
			"image_name":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"internal_ips":schema.ListAttribute{
				ElementType: types.StringType,
				Computed:            true,
				MarkdownDescription: "",
			},
			"is_limit_based":schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"is_usage_based":schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"key_fingerprint":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"key_name":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"latitude":schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"longitude":schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_category_name":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_category_uuid":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_offering_name":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_offering_uuid":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_plan_uuid":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_resource_state":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"marketplace_resource_uuid":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"min_disk":schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Minimum disk size in MiB",
			},
			"min_ram":schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Minimum memory size in MiB",
			},
			"modified":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"ports":schema.ListAttribute{
				// Flatten complex objects to list of strings for simpler schema
				ElementType: types.StringType,
				Computed:            true,
				MarkdownDescription: "",
			},
			"ram":schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Memory size in MiB",
			},
			"resource_type":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"security_groups":schema.ListAttribute{
				// Flatten complex objects to list of strings for simpler schema
				ElementType: types.StringType,
				Computed:            true,
				MarkdownDescription: "",
			},
			"server_group":schema.ObjectAttribute{
				Computed: true,
				AttributeTypes: map[string]attr.Type{
					"name": types.StringType,
					"state": types.StringType,
					"url": types.StringType,
				},
				MarkdownDescription: "",
			},
			"service_name":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_settings":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "OpenStack provider settings",
			},
			"service_settings_error_message":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"service_settings_state":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"start_time":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"url":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"user_data":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Additional data that will be added to instance on provisioning",
			},
			"volumes":schema.ListAttribute{
				// Flatten complex objects to list of strings for simpler schema
				ElementType: types.StringType,
				Computed:            true,
				MarkdownDescription: "List of volumes attached to the instance",
			},
		},
	}
}

func (d *OpenstackInstanceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}

func (d *OpenstackInstanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OpenstackInstanceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if UUID is provided for direct lookup
	if !data.UUID.IsNull() && data.UUID.ValueString() != "" {
		var item map[string]interface{}
		
		err := d.client.GetByUUID(ctx, "/api/openstack-instances/{uuid}/", data.UUID.ValueString(), &item)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Read Instance",
				"An error occurred while reading the openstack_instance by UUID: "+err.Error(),
			)
			return
		}

		// Extract data from single result
		if uuid, ok := item["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}
		if val, ok := item["attach_volume_uuid"].(string); ok {
			data.AttachVolumeUuid = types.StringValue(val)
		}
		if val, ok := item["availability_zone_name"].(string); ok {
			data.AvailabilityZoneName = types.StringValue(val)
		}
		if val, ok := item["backend_id"].(string); ok {
			data.BackendId = types.StringValue(val)
		}
		if val, ok := item["customer"].(string); ok {
			data.Customer = types.StringValue(val)
		}
		if val, ok := item["customer_abbreviation"].(string); ok {
			data.CustomerAbbreviation = types.StringValue(val)
		}
		if val, ok := item["customer_name"].(string); ok {
			data.CustomerName = types.StringValue(val)
		}
		if val, ok := item["customer_native_name"].(string); ok {
			data.CustomerNativeName = types.StringValue(val)
		}
		if val, ok := item["customer_uuid"].(string); ok {
			data.CustomerUuid = types.StringValue(val)
		}
		if val, ok := item["description"].(string); ok {
			data.Description = types.StringValue(val)
		}
		if val, ok := item["external_ip"].(string); ok {
			data.ExternalIp = types.StringValue(val)
		}
		if val, ok := item["field"].(string); ok {
			data.Field = types.StringValue(val)
		}
		if val, ok := item["name"].(string); ok {
			data.Name = types.StringValue(val)
		}
		if val, ok := item["name_exact"].(string); ok {
			data.NameExact = types.StringValue(val)
		}
		if val, ok := item["project"].(string); ok {
			data.Project = types.StringValue(val)
		}
		if val, ok := item["project_name"].(string); ok {
			data.ProjectName = types.StringValue(val)
		}
		if val, ok := item["project_uuid"].(string); ok {
			data.ProjectUuid = types.StringValue(val)
		}
		if val, ok := item["query"].(string); ok {
			data.Query = types.StringValue(val)
		}
		if val, ok := item["runtime_state"].(string); ok {
			data.RuntimeState = types.StringValue(val)
		}
		if val, ok := item["service_settings_name"].(string); ok {
			data.ServiceSettingsName = types.StringValue(val)
		}
		if val, ok := item["service_settings_uuid"].(string); ok {
			data.ServiceSettingsUuid = types.StringValue(val)
		}
		if val, ok := item["state"].(string); ok {
			data.State = types.StringValue(val)
		}
		if val, ok := item["tenant"].(string); ok {
			data.Tenant = types.StringValue(val)
		}
		if val, ok := item["tenant_uuid"].(string); ok {
			data.TenantUuid = types.StringValue(val)
		}
		if val, ok := item["uuid"].(string); ok {
			data.Uuid = types.StringValue(val)
		}

	sourceMap := item
	// Map response fields to data model
	if val, ok := sourceMap["access_url"]; ok {
		if val == nil {
			data.AccessUrl = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.AccessUrl = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["action"]; ok {
		if val == nil {
			data.Action = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Action = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["availability_zone"]; ok {
		if val == nil {
			data.AvailabilityZone = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.AvailabilityZone = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["connect_directly_to_external_network"]; ok {
		if val == nil {
			data.ConnectDirectlyToExternalNetwork = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.ConnectDirectlyToExternalNetwork = types.BoolValue(b)
		}
		}
	}
	if val, ok := sourceMap["cores"]; ok {
		if val == nil {
			data.Cores = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.Cores = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["created"]; ok {
		if val == nil {
			data.Created = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["disk"]; ok {
		if val == nil {
			data.Disk = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.Disk = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["error_message"]; ok {
		if val == nil {
			data.ErrorMessage = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ErrorMessage = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["error_traceback"]; ok {
		if val == nil {
			data.ErrorTraceback = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ErrorTraceback = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["external_address"]; ok {
		if val == nil {
			data.ExternalAddress = types.ListNull(types.StringType)
		} else {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.ExternalAddress = listVal
		}
		}
	}
	if val, ok := sourceMap["external_ips"]; ok {
		if val == nil {
			data.ExternalIps = types.ListNull(types.StringType)
		} else {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.ExternalIps = listVal
		}
		}
	}
	if val, ok := sourceMap["flavor_disk"]; ok {
		if val == nil {
			data.FlavorDisk = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.FlavorDisk = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["flavor_name"]; ok {
		if val == nil {
			data.FlavorName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.FlavorName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["floating_ips"]; ok {
		if val == nil {
			data.FloatingIps = types.ListNull(types.StringType)
		} else {
		// List of objects - flatten to list of URLs/UUIDs to avoid schema mismatch
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if obj, ok := item.(map[string]interface{}); ok {
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.FloatingIps = listVal
		} else {
			data.FloatingIps = types.ListNull(types.StringType)
		}
		}
	}
	if val, ok := sourceMap["hypervisor_hostname"]; ok {
		if val == nil {
			data.HypervisorHostname = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.HypervisorHostname = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["image_name"]; ok {
		if val == nil {
			data.ImageName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ImageName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["internal_ips"]; ok {
		if val == nil {
			data.InternalIps = types.ListNull(types.StringType)
		} else {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.InternalIps = listVal
		}
		}
	}
	if val, ok := sourceMap["is_limit_based"]; ok {
		if val == nil {
			data.IsLimitBased = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.IsLimitBased = types.BoolValue(b)
		}
		}
	}
	if val, ok := sourceMap["is_usage_based"]; ok {
		if val == nil {
			data.IsUsageBased = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.IsUsageBased = types.BoolValue(b)
		}
		}
	}
	if val, ok := sourceMap["key_fingerprint"]; ok {
		if val == nil {
			data.KeyFingerprint = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.KeyFingerprint = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["key_name"]; ok {
		if val == nil {
			data.KeyName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.KeyName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["latitude"]; ok {
		if val == nil {
			data.Latitude = types.Float64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.Latitude = types.Float64Value(num)
		}
		}
	}
	if val, ok := sourceMap["longitude"]; ok {
		if val == nil {
			data.Longitude = types.Float64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.Longitude = types.Float64Value(num)
		}
		}
	}
	if val, ok := sourceMap["marketplace_category_name"]; ok {
		if val == nil {
			data.MarketplaceCategoryName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplaceCategoryName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["marketplace_category_uuid"]; ok {
		if val == nil {
			data.MarketplaceCategoryUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplaceCategoryUuid = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["marketplace_offering_name"]; ok {
		if val == nil {
			data.MarketplaceOfferingName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplaceOfferingName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["marketplace_offering_uuid"]; ok {
		if val == nil {
			data.MarketplaceOfferingUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplaceOfferingUuid = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["marketplace_plan_uuid"]; ok {
		if val == nil {
			data.MarketplacePlanUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplacePlanUuid = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["marketplace_resource_state"]; ok {
		if val == nil {
			data.MarketplaceResourceState = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplaceResourceState = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["marketplace_resource_uuid"]; ok {
		if val == nil {
			data.MarketplaceResourceUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplaceResourceUuid = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["min_disk"]; ok {
		if val == nil {
			data.MinDisk = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.MinDisk = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["min_ram"]; ok {
		if val == nil {
			data.MinRam = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.MinRam = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["modified"]; ok {
		if val == nil {
			data.Modified = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Modified = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["ports"]; ok {
		if val == nil {
			data.Ports = types.ListNull(types.StringType)
		} else {
		// List of objects - flatten to list of URLs/UUIDs to avoid schema mismatch
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if obj, ok := item.(map[string]interface{}); ok {
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.Ports = listVal
		} else {
			data.Ports = types.ListNull(types.StringType)
		}
		}
	}
	if val, ok := sourceMap["ram"]; ok {
		if val == nil {
			data.Ram = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.Ram = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["resource_type"]; ok {
		if val == nil {
			data.ResourceType = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ResourceType = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["security_groups"]; ok {
		if val == nil {
			data.SecurityGroups = types.ListNull(types.StringType)
		} else {
		// List of objects - flatten to list of URLs/UUIDs to avoid schema mismatch
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if obj, ok := item.(map[string]interface{}); ok {
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.SecurityGroups = listVal
		} else {
			data.SecurityGroups = types.ListNull(types.StringType)
		}
		}
	}
	if val, ok := sourceMap["server_group"]; ok {
		if val == nil {
			data.ServerGroup = types.ObjectNull(map[string]attr.Type{
				"name": types.StringType,
				"state": types.StringType,
				"url": types.StringType,
			})
		} else {
		// Nested object
		if objMap, ok := val.(map[string]interface{}); ok {
			attrTypes := map[string]attr.Type{
				"name": types.StringType,
				"state": types.StringType,
				"url": types.StringType,
			}
			attrValues := map[string]attr.Value{
				"name": types.StringValue(getStringFromMap(objMap, "name")),
				"state": types.StringValue(getStringFromMap(objMap, "state")),
				"url": types.StringValue(getStringFromMap(objMap, "url")),
			}
			objVal, _ := types.ObjectValue(attrTypes, attrValues)
			data.ServerGroup = objVal
		}
		}
	}
	if val, ok := sourceMap["service_name"]; ok {
		if val == nil {
			data.ServiceName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ServiceName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["service_settings"]; ok {
		if val == nil {
			data.ServiceSettings = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ServiceSettings = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["service_settings_error_message"]; ok {
		if val == nil {
			data.ServiceSettingsErrorMessage = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ServiceSettingsErrorMessage = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["service_settings_state"]; ok {
		if val == nil {
			data.ServiceSettingsState = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ServiceSettingsState = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["start_time"]; ok {
		if val == nil {
			data.StartTime = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.StartTime = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["url"]; ok {
		if val == nil {
			data.Url = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["user_data"]; ok {
		if val == nil {
			data.UserData = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.UserData = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["volumes"]; ok {
		if val == nil {
			data.Volumes = types.ListNull(types.StringType)
		} else {
		// List of objects - flatten to list of URLs/UUIDs to avoid schema mismatch
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if obj, ok := item.(map[string]interface{}); ok {
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.Volumes = listVal
		} else {
			data.Volumes = types.ListNull(types.StringType)
		}
		}
	}

	} else {
		// Filter by provided parameters
		var results []map[string]interface{}
		
		filters := map[string]string{}
		if !data.AttachVolumeUuid.IsNull() {
			filters["attach_volume_uuid"] = data.AttachVolumeUuid.ValueString()
		}
		if !data.AvailabilityZoneName.IsNull() {
			filters["availability_zone_name"] = data.AvailabilityZoneName.ValueString()
		}
		if !data.BackendId.IsNull() {
			filters["backend_id"] = data.BackendId.ValueString()
		}
		if !data.CanManage.IsNull() {
			filters["can_manage"] = fmt.Sprintf("%t", data.CanManage.ValueBool())
		}
		if !data.Customer.IsNull() {
			filters["customer"] = data.Customer.ValueString()
		}
		if !data.CustomerAbbreviation.IsNull() {
			filters["customer_abbreviation"] = data.CustomerAbbreviation.ValueString()
		}
		if !data.CustomerName.IsNull() {
			filters["customer_name"] = data.CustomerName.ValueString()
		}
		if !data.CustomerNativeName.IsNull() {
			filters["customer_native_name"] = data.CustomerNativeName.ValueString()
		}
		if !data.CustomerUuid.IsNull() {
			filters["customer_uuid"] = data.CustomerUuid.ValueString()
		}
		if !data.Description.IsNull() {
			filters["description"] = data.Description.ValueString()
		}
		if !data.ExternalIp.IsNull() {
			filters["external_ip"] = data.ExternalIp.ValueString()
		}
		if !data.Field.IsNull() {
			filters["field"] = data.Field.ValueString()
		}
		if !data.Name.IsNull() {
			filters["name"] = data.Name.ValueString()
		}
		if !data.NameExact.IsNull() {
			filters["name_exact"] = data.NameExact.ValueString()
		}
		if !data.Page.IsNull() {
			filters["page"] = fmt.Sprintf("%d", data.Page.ValueInt64())
		}
		if !data.PageSize.IsNull() {
			filters["page_size"] = fmt.Sprintf("%d", data.PageSize.ValueInt64())
		}
		if !data.Project.IsNull() {
			filters["project"] = data.Project.ValueString()
		}
		if !data.ProjectName.IsNull() {
			filters["project_name"] = data.ProjectName.ValueString()
		}
		if !data.ProjectUuid.IsNull() {
			filters["project_uuid"] = data.ProjectUuid.ValueString()
		}
		if !data.Query.IsNull() {
			filters["query"] = data.Query.ValueString()
		}
		if !data.RuntimeState.IsNull() {
			filters["runtime_state"] = data.RuntimeState.ValueString()
		}
		if !data.ServiceSettingsName.IsNull() {
			filters["service_settings_name"] = data.ServiceSettingsName.ValueString()
		}
		if !data.ServiceSettingsUuid.IsNull() {
			filters["service_settings_uuid"] = data.ServiceSettingsUuid.ValueString()
		}
		if !data.State.IsNull() {
			filters["state"] = data.State.ValueString()
		}
		if !data.Tenant.IsNull() {
			filters["tenant"] = data.Tenant.ValueString()
		}
		if !data.TenantUuid.IsNull() {
			filters["tenant_uuid"] = data.TenantUuid.ValueString()
		}
		if !data.Uuid.IsNull() {
			filters["uuid"] = data.Uuid.ValueString()
		}
		
		if len(filters) == 0 {
			resp.Diagnostics.AddError(
				"Missing Filter Parameters",
				"At least one filter parameter (or 'id') must be provided to lookup openstack_instance.",
			)
			return
		}
		
		err := d.client.ListWithFilter(ctx, "/api/openstack-instances/", filters, &results)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List Instance",
				"An error occurred while filtering openstack_instance: "+err.Error(),
			)
			return
		}

		// Check results
		if len(results) == 0 {
			resp.Diagnostics.AddError(
				"Instance Not Found",
				"No openstack_instance found with provided filters.",
			)
			return
		}
		
		if len(results) > 1 {
			resp.Diagnostics.AddError(
				"Multiple Instances Found",
				fmt.Sprintf("Found %d openstack_instances with provided filters. Please use more specific filters or lookup by UUID.", len(results)),
			)
			return
		}

		// Extract data from the single result
		if uuid, ok := results[0]["uuid"].(string); ok {
			data.UUID = types.StringValue(uuid)
		}
		if val, ok := results[0]["attach_volume_uuid"].(string); ok {
			data.AttachVolumeUuid = types.StringValue(val)
		}
		if val, ok := results[0]["availability_zone_name"].(string); ok {
			data.AvailabilityZoneName = types.StringValue(val)
		}
		if val, ok := results[0]["backend_id"].(string); ok {
			data.BackendId = types.StringValue(val)
		}
		if val, ok := results[0]["customer"].(string); ok {
			data.Customer = types.StringValue(val)
		}
		if val, ok := results[0]["customer_abbreviation"].(string); ok {
			data.CustomerAbbreviation = types.StringValue(val)
		}
		if val, ok := results[0]["customer_name"].(string); ok {
			data.CustomerName = types.StringValue(val)
		}
		if val, ok := results[0]["customer_native_name"].(string); ok {
			data.CustomerNativeName = types.StringValue(val)
		}
		if val, ok := results[0]["customer_uuid"].(string); ok {
			data.CustomerUuid = types.StringValue(val)
		}
		if val, ok := results[0]["description"].(string); ok {
			data.Description = types.StringValue(val)
		}
		if val, ok := results[0]["external_ip"].(string); ok {
			data.ExternalIp = types.StringValue(val)
		}
		if val, ok := results[0]["field"].(string); ok {
			data.Field = types.StringValue(val)
		}
		if val, ok := results[0]["name"].(string); ok {
			data.Name = types.StringValue(val)
		}
		if val, ok := results[0]["name_exact"].(string); ok {
			data.NameExact = types.StringValue(val)
		}
		if val, ok := results[0]["project"].(string); ok {
			data.Project = types.StringValue(val)
		}
		if val, ok := results[0]["project_name"].(string); ok {
			data.ProjectName = types.StringValue(val)
		}
		if val, ok := results[0]["project_uuid"].(string); ok {
			data.ProjectUuid = types.StringValue(val)
		}
		if val, ok := results[0]["query"].(string); ok {
			data.Query = types.StringValue(val)
		}
		if val, ok := results[0]["runtime_state"].(string); ok {
			data.RuntimeState = types.StringValue(val)
		}
		if val, ok := results[0]["service_settings_name"].(string); ok {
			data.ServiceSettingsName = types.StringValue(val)
		}
		if val, ok := results[0]["service_settings_uuid"].(string); ok {
			data.ServiceSettingsUuid = types.StringValue(val)
		}
		if val, ok := results[0]["state"].(string); ok {
			data.State = types.StringValue(val)
		}
		if val, ok := results[0]["tenant"].(string); ok {
			data.Tenant = types.StringValue(val)
		}
		if val, ok := results[0]["tenant_uuid"].(string); ok {
			data.TenantUuid = types.StringValue(val)
		}
		if val, ok := results[0]["uuid"].(string); ok {
			data.Uuid = types.StringValue(val)
		}
	sourceMap := results[0]
	// Map response fields to data model
	if val, ok := sourceMap["access_url"]; ok {
		if val == nil {
			data.AccessUrl = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.AccessUrl = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["action"]; ok {
		if val == nil {
			data.Action = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Action = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["availability_zone"]; ok {
		if val == nil {
			data.AvailabilityZone = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.AvailabilityZone = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["connect_directly_to_external_network"]; ok {
		if val == nil {
			data.ConnectDirectlyToExternalNetwork = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.ConnectDirectlyToExternalNetwork = types.BoolValue(b)
		}
		}
	}
	if val, ok := sourceMap["cores"]; ok {
		if val == nil {
			data.Cores = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.Cores = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["created"]; ok {
		if val == nil {
			data.Created = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Created = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["disk"]; ok {
		if val == nil {
			data.Disk = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.Disk = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["error_message"]; ok {
		if val == nil {
			data.ErrorMessage = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ErrorMessage = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["error_traceback"]; ok {
		if val == nil {
			data.ErrorTraceback = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ErrorTraceback = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["external_address"]; ok {
		if val == nil {
			data.ExternalAddress = types.ListNull(types.StringType)
		} else {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.ExternalAddress = listVal
		}
		}
	}
	if val, ok := sourceMap["external_ips"]; ok {
		if val == nil {
			data.ExternalIps = types.ListNull(types.StringType)
		} else {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.ExternalIps = listVal
		}
		}
	}
	if val, ok := sourceMap["flavor_disk"]; ok {
		if val == nil {
			data.FlavorDisk = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.FlavorDisk = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["flavor_name"]; ok {
		if val == nil {
			data.FlavorName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.FlavorName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["floating_ips"]; ok {
		if val == nil {
			data.FloatingIps = types.ListNull(types.StringType)
		} else {
		// List of objects - flatten to list of URLs/UUIDs to avoid schema mismatch
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if obj, ok := item.(map[string]interface{}); ok {
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.FloatingIps = listVal
		} else {
			data.FloatingIps = types.ListNull(types.StringType)
		}
		}
	}
	if val, ok := sourceMap["hypervisor_hostname"]; ok {
		if val == nil {
			data.HypervisorHostname = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.HypervisorHostname = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["image_name"]; ok {
		if val == nil {
			data.ImageName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ImageName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["internal_ips"]; ok {
		if val == nil {
			data.InternalIps = types.ListNull(types.StringType)
		} else {
		// List of strings (or flattened objects)
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if str, ok := item.(string); ok {
					items = append(items, types.StringValue(str))
				} else if obj, ok := item.(map[string]interface{}); ok {
					// Flattening logic: extract URL or UUID
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					} else if name, ok := obj["name"].(string); ok {
						items = append(items, types.StringValue(name))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.InternalIps = listVal
		}
		}
	}
	if val, ok := sourceMap["is_limit_based"]; ok {
		if val == nil {
			data.IsLimitBased = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.IsLimitBased = types.BoolValue(b)
		}
		}
	}
	if val, ok := sourceMap["is_usage_based"]; ok {
		if val == nil {
			data.IsUsageBased = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.IsUsageBased = types.BoolValue(b)
		}
		}
	}
	if val, ok := sourceMap["key_fingerprint"]; ok {
		if val == nil {
			data.KeyFingerprint = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.KeyFingerprint = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["key_name"]; ok {
		if val == nil {
			data.KeyName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.KeyName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["latitude"]; ok {
		if val == nil {
			data.Latitude = types.Float64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.Latitude = types.Float64Value(num)
		}
		}
	}
	if val, ok := sourceMap["longitude"]; ok {
		if val == nil {
			data.Longitude = types.Float64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.Longitude = types.Float64Value(num)
		}
		}
	}
	if val, ok := sourceMap["marketplace_category_name"]; ok {
		if val == nil {
			data.MarketplaceCategoryName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplaceCategoryName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["marketplace_category_uuid"]; ok {
		if val == nil {
			data.MarketplaceCategoryUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplaceCategoryUuid = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["marketplace_offering_name"]; ok {
		if val == nil {
			data.MarketplaceOfferingName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplaceOfferingName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["marketplace_offering_uuid"]; ok {
		if val == nil {
			data.MarketplaceOfferingUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplaceOfferingUuid = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["marketplace_plan_uuid"]; ok {
		if val == nil {
			data.MarketplacePlanUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplacePlanUuid = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["marketplace_resource_state"]; ok {
		if val == nil {
			data.MarketplaceResourceState = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplaceResourceState = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["marketplace_resource_uuid"]; ok {
		if val == nil {
			data.MarketplaceResourceUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.MarketplaceResourceUuid = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["min_disk"]; ok {
		if val == nil {
			data.MinDisk = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.MinDisk = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["min_ram"]; ok {
		if val == nil {
			data.MinRam = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.MinRam = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["modified"]; ok {
		if val == nil {
			data.Modified = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Modified = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["ports"]; ok {
		if val == nil {
			data.Ports = types.ListNull(types.StringType)
		} else {
		// List of objects - flatten to list of URLs/UUIDs to avoid schema mismatch
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if obj, ok := item.(map[string]interface{}); ok {
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.Ports = listVal
		} else {
			data.Ports = types.ListNull(types.StringType)
		}
		}
	}
	if val, ok := sourceMap["ram"]; ok {
		if val == nil {
			data.Ram = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.Ram = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["resource_type"]; ok {
		if val == nil {
			data.ResourceType = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ResourceType = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["security_groups"]; ok {
		if val == nil {
			data.SecurityGroups = types.ListNull(types.StringType)
		} else {
		// List of objects - flatten to list of URLs/UUIDs to avoid schema mismatch
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if obj, ok := item.(map[string]interface{}); ok {
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.SecurityGroups = listVal
		} else {
			data.SecurityGroups = types.ListNull(types.StringType)
		}
		}
	}
	if val, ok := sourceMap["server_group"]; ok {
		if val == nil {
			data.ServerGroup = types.ObjectNull(map[string]attr.Type{
				"name": types.StringType,
				"state": types.StringType,
				"url": types.StringType,
			})
		} else {
		// Nested object
		if objMap, ok := val.(map[string]interface{}); ok {
			attrTypes := map[string]attr.Type{
				"name": types.StringType,
				"state": types.StringType,
				"url": types.StringType,
			}
			attrValues := map[string]attr.Value{
				"name": types.StringValue(getStringFromMap(objMap, "name")),
				"state": types.StringValue(getStringFromMap(objMap, "state")),
				"url": types.StringValue(getStringFromMap(objMap, "url")),
			}
			objVal, _ := types.ObjectValue(attrTypes, attrValues)
			data.ServerGroup = objVal
		}
		}
	}
	if val, ok := sourceMap["service_name"]; ok {
		if val == nil {
			data.ServiceName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ServiceName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["service_settings"]; ok {
		if val == nil {
			data.ServiceSettings = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ServiceSettings = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["service_settings_error_message"]; ok {
		if val == nil {
			data.ServiceSettingsErrorMessage = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ServiceSettingsErrorMessage = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["service_settings_state"]; ok {
		if val == nil {
			data.ServiceSettingsState = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ServiceSettingsState = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["start_time"]; ok {
		if val == nil {
			data.StartTime = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.StartTime = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["url"]; ok {
		if val == nil {
			data.Url = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Url = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["user_data"]; ok {
		if val == nil {
			data.UserData = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.UserData = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["volumes"]; ok {
		if val == nil {
			data.Volumes = types.ListNull(types.StringType)
		} else {
		// List of objects - flatten to list of URLs/UUIDs to avoid schema mismatch
		if arr, ok := val.([]interface{}); ok {
			items := make([]attr.Value, 0, len(arr))
			for _, item := range arr {
				if obj, ok := item.(map[string]interface{}); ok {
					if url, ok := obj["url"].(string); ok {
						parts := strings.Split(strings.TrimRight(url, "/"), "/")
						uuid := parts[len(parts)-1]
						items = append(items, types.StringValue(uuid))
					} else if uuid, ok := obj["uuid"].(string); ok {
						items = append(items, types.StringValue(uuid))
					}
				}
			}
			listVal, _ := types.ListValue(types.StringType, items)
			data.Volumes = listVal
		} else {
			data.Volumes = types.ListNull(types.StringType)
		}
		}
	}
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
