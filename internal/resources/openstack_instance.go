package resources

import (
	"context"
    "strings"
    "fmt"
    "time"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-waldur-provider/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OpenstackInstanceResource{}
var _ resource.ResourceWithImportState = &OpenstackInstanceResource{}

func NewOpenstackInstanceResource() resource.Resource {
	return &OpenstackInstanceResource{}
}

// OpenstackInstanceResource defines the resource implementation.
type OpenstackInstanceResource struct {
	client *client.Client
}

// OpenstackInstanceResourceModel describes the resource data model.
type OpenstackInstanceResourceModel struct {
	UUID types.String `tfsdk:"id"`
	AvailabilityZone types.String `tfsdk:"availability_zone"`
	ConnectDirectlyToExternalNetwork types.Bool `tfsdk:"connect_directly_to_external_network"`
	DataVolumeSize types.Int64 `tfsdk:"data_volume_size"`
	DataVolumeType types.String `tfsdk:"data_volume_type"`
	DataVolumes types.List `tfsdk:"data_volumes"`
	Description types.String `tfsdk:"description"`
	Flavor types.String `tfsdk:"flavor"`
	FloatingIps types.List `tfsdk:"floating_ips"`
	Image types.String `tfsdk:"image"`
	Name types.String `tfsdk:"name"`
	Ports types.List `tfsdk:"ports"`
	SecurityGroups types.List `tfsdk:"security_groups"`
	SshPublicKey types.String `tfsdk:"ssh_public_key"`
	SystemVolumeSize types.Int64 `tfsdk:"system_volume_size"`
	SystemVolumeType types.String `tfsdk:"system_volume_type"`
	UserData types.String `tfsdk:"user_data"`
	Project types.String `tfsdk:"project"`
	Offering types.String `tfsdk:"offering"`
	AccessUrl types.String `tfsdk:"access_url"`
	Action types.String `tfsdk:"action"`
	AvailabilityZoneName types.String `tfsdk:"availability_zone_name"`
	BackendId types.String `tfsdk:"backend_id"`
	Cores types.Int64 `tfsdk:"cores"`
	Created types.String `tfsdk:"created"`
	Customer types.String `tfsdk:"customer"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerName types.String `tfsdk:"customer_name"`
	CustomerNativeName types.String `tfsdk:"customer_native_name"`
	CustomerUuid types.String `tfsdk:"customer_uuid"`
	Disk types.Int64 `tfsdk:"disk"`
	ErrorMessage types.String `tfsdk:"error_message"`
	ErrorTraceback types.String `tfsdk:"error_traceback"`
	ExternalAddress types.List `tfsdk:"external_address"`
	ExternalIps types.List `tfsdk:"external_ips"`
	FlavorDisk types.Int64 `tfsdk:"flavor_disk"`
	FlavorName types.String `tfsdk:"flavor_name"`
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
	ProjectName types.String `tfsdk:"project_name"`
	ProjectUuid types.String `tfsdk:"project_uuid"`
	Ram types.Int64 `tfsdk:"ram"`
	ResourceType types.String `tfsdk:"resource_type"`
	RuntimeState types.String `tfsdk:"runtime_state"`
	ServerGroup types.Object `tfsdk:"server_group"`
	ServiceName types.String `tfsdk:"service_name"`
	ServiceSettings types.String `tfsdk:"service_settings"`
	ServiceSettingsErrorMessage types.String `tfsdk:"service_settings_error_message"`
	ServiceSettingsState types.String `tfsdk:"service_settings_state"`
	ServiceSettingsUuid types.String `tfsdk:"service_settings_uuid"`
	StartTime types.String `tfsdk:"start_time"`
	Tenant types.String `tfsdk:"tenant"`
	TenantUuid types.String `tfsdk:"tenant_uuid"`
	Url types.String `tfsdk:"url"`
	Volumes types.List `tfsdk:"volumes"`
}

func (r *OpenstackInstanceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_openstack_instance"
}

func (r *OpenstackInstanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenstackInstance resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"availability_zone":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Availability zone where this instance is located",
			},
			"connect_directly_to_external_network":schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "If True, instance will be connected directly to external network",
			},
			"data_volume_size":schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Size of the data volume in MiB. Minimum size is 1024 MiB (1 GiB)",
			},
			"data_volume_type":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Volume type for the data volume",
			},
			"data_volumes":schema.ListAttribute{
				// Flatten complex objects to list of UUIDs for simpler schema
				ElementType: types.StringType,
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Additional data volumes to attach to the instance",
			},
			"description":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"flavor":schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The flavor to use for the instance",
			},
			"floating_ips":schema.ListAttribute{
				// Flatten complex objects to list of UUIDs for simpler schema
				ElementType: types.StringType,
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Floating IPs to assign to the instance",
			},
			"image":schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The OS image to use for the instance",
			},
			"name":schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "",
			},
			"ports":schema.ListAttribute{
				// Flatten complex objects to list of UUIDs for simpler schema
				ElementType: types.StringType,
				Required:            true,
				MarkdownDescription: "Network ports to attach to the instance",
			},
			"security_groups":schema.ListAttribute{
				// Flatten complex objects to list of UUIDs for simpler schema
				ElementType: types.StringType,
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "List of security groups to apply to the instance",
			},
			"ssh_public_key":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"system_volume_size":schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Size of the system volume in MiB. Minimum size is 1024 MiB (1 GiB)",
			},
			"system_volume_type":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Volume type for the system volume",
			},
			"user_data":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Additional data that will be added to instance on provisioning",
			},
			"project":schema.StringAttribute{
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Project UUID",
			},
			"offering":schema.StringAttribute{
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Offering UUID or URL",
			},
			"access_url":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"action":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"availability_zone_name":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Name of the availability zone where instance is located",
			},
			"backend_id":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Instance ID in the OpenStack backend",
			},
			"cores":schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of cores in a VM",
			},
			"created":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_abbreviation":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_name":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_native_name":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_uuid":schema.StringAttribute{
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
			"project_name":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"project_uuid":schema.StringAttribute{
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
			"runtime_state":schema.StringAttribute{
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
			"service_settings_uuid":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"start_time":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"tenant":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The OpenStack tenant to create the instance in",
			},
			"tenant_uuid":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the OpenStack tenant",
			},
			"url":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"volumes":schema.ListAttribute{
				// Flatten complex objects to list of UUIDs for simpler schema
				ElementType: types.StringType,
				Computed:            true,
				MarkdownDescription: "List of volumes attached to the instance",
			},
		},
	}
}

func (r *OpenstackInstanceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *client.Client, got something else. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = client
}

func (r *OpenstackInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpenstackInstanceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Phase 1: Payload Construction
	attributes := map[string]interface{}{}
	if !data.AvailabilityZone.IsNull() && !data.AvailabilityZone.IsUnknown() && data.AvailabilityZone.ValueString() != "" {
		attributes["availability_zone"] = data.AvailabilityZone.ValueString()
	}
	if !data.ConnectDirectlyToExternalNetwork.IsNull() && !data.ConnectDirectlyToExternalNetwork.IsUnknown() {
		attributes["connect_directly_to_external_network"] = data.ConnectDirectlyToExternalNetwork.ValueBool()
	}
	if !data.DataVolumeSize.IsNull() && !data.DataVolumeSize.IsUnknown() && data.DataVolumeSize.ValueInt64() != 0 {
		attributes["data_volume_size"] = data.DataVolumeSize.ValueInt64()
	}
	if !data.DataVolumeType.IsNull() && !data.DataVolumeType.IsUnknown() && data.DataVolumeType.ValueString() != "" {
		attributes["data_volume_type"] = data.DataVolumeType.ValueString()
	}
	if !data.DataVolumes.IsNull() && !data.DataVolumes.IsUnknown() && len(data.DataVolumes.Elements()) > 0 {
		var items []interface{}
		for _, elem := range data.DataVolumes.Elements() {
			// Schema flattened to strings - transform back to objects for API
			if strVal, ok := elem.(types.String); ok && !strVal.IsNull() && strVal.ValueString() != "" {
				// Default: use "url" as the key
				items = append(items, map[string]interface{}{"url": strVal.ValueString()})
			}
		}
		if len(items) > 0 {
			attributes["data_volumes"] = items
		}
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() && data.Description.ValueString() != "" {
		attributes["description"] = data.Description.ValueString()
	}
	if !data.Flavor.IsNull() && !data.Flavor.IsUnknown() && data.Flavor.ValueString() != "" {
		attributes["flavor"] = data.Flavor.ValueString()
	}
	if !data.FloatingIps.IsNull() && !data.FloatingIps.IsUnknown() && len(data.FloatingIps.Elements()) > 0 {
		var items []interface{}
		for _, elem := range data.FloatingIps.Elements() {
			// Schema flattened to strings - transform back to objects for API
			if strVal, ok := elem.(types.String); ok && !strVal.IsNull() && strVal.ValueString() != "" {
				// floating_ips expects {"subnet": "url"} or {"url": "url"}
				items = append(items, map[string]interface{}{"url": strVal.ValueString()})
			}
		}
		if len(items) > 0 {
			attributes["floating_ips"] = items
		}
	}
	if !data.Image.IsNull() && !data.Image.IsUnknown() && data.Image.ValueString() != "" {
		attributes["image"] = data.Image.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() && data.Name.ValueString() != "" {
		attributes["name"] = data.Name.ValueString()
	}
	if !data.Ports.IsNull() && !data.Ports.IsUnknown() && len(data.Ports.Elements()) > 0 {
		var items []interface{}
		for _, elem := range data.Ports.Elements() {
			// Schema flattened to strings - transform back to objects for API
			if strVal, ok := elem.(types.String); ok && !strVal.IsNull() && strVal.ValueString() != "" {
				// ports expects {"subnet": "url"}
				items = append(items, map[string]interface{}{"subnet": strVal.ValueString()})
			}
		}
		if len(items) > 0 {
			attributes["ports"] = items
		}
	}
	if !data.SecurityGroups.IsNull() && !data.SecurityGroups.IsUnknown() && len(data.SecurityGroups.Elements()) > 0 {
		var items []interface{}
		for _, elem := range data.SecurityGroups.Elements() {
			// Schema flattened to strings - transform back to objects for API
			if strVal, ok := elem.(types.String); ok && !strVal.IsNull() && strVal.ValueString() != "" {
				// security_groups expects {"url": "url"}
				items = append(items, map[string]interface{}{"url": strVal.ValueString()})
			}
		}
		if len(items) > 0 {
			attributes["security_groups"] = items
		}
	}
	if !data.SshPublicKey.IsNull() && !data.SshPublicKey.IsUnknown() && data.SshPublicKey.ValueString() != "" {
		attributes["ssh_public_key"] = data.SshPublicKey.ValueString()
	}
	if !data.SystemVolumeSize.IsNull() && !data.SystemVolumeSize.IsUnknown() && data.SystemVolumeSize.ValueInt64() != 0 {
		attributes["system_volume_size"] = data.SystemVolumeSize.ValueInt64()
	}
	if !data.SystemVolumeType.IsNull() && !data.SystemVolumeType.IsUnknown() && data.SystemVolumeType.ValueString() != "" {
		attributes["system_volume_type"] = data.SystemVolumeType.ValueString()
	}
	if !data.UserData.IsNull() && !data.UserData.IsUnknown() && data.UserData.ValueString() != "" {
		attributes["user_data"] = data.UserData.ValueString()
	}

	// Format offering URL - must use marketplace-public-offerings endpoint
	offeringValue := data.Offering.ValueString()
	var offeringURL string
	if strings.HasPrefix(offeringValue, "http") {
		// Already a URL - ensure it uses marketplace-public-offerings
		offeringURL = strings.Replace(offeringValue, "marketplace-provider-offerings", "marketplace-public-offerings", 1)
	} else {
		// UUID only - construct full URL
		offeringURL = fmt.Sprintf("/api/marketplace-public-offerings/%s/", offeringValue)
	}

	payload := map[string]interface{}{
		"project": fmt.Sprintf("/api/projects/%s/", data.Project.ValueString()),
		"offering": offeringURL,
		"attributes": attributes,
	}

	// Phase 2: Submit Order
	var orderRes map[string]interface{}
	err := r.client.Post(ctx, "/api/marketplace-orders/", payload, &orderRes)
	if err != nil {
		resp.Diagnostics.AddError("Order Submission Failed", err.Error())
		return
	}
	
	orderUUID, ok := orderRes["uuid"].(string)
	if !ok {
		resp.Diagnostics.AddError("Invalid Response", "Order UUID not found")
		return
	}

	// Phase 3: Poll for Completion
	stateConf := &retry.StateChangeConf{
		Pending: []string{"pending", "executing", "created"},
		Target:  []string{"done"},
		Refresh: func() (interface{}, string, error) {
			var res map[string]interface{}
			err := r.client.GetByUUID(ctx, "/api/marketplace-orders/{uuid}/", orderUUID, &res)
			if err != nil {
				return nil, "", err
			}
			
			state, _ := res["state"].(string)
			if state == "erred" || state == "rejected" {
				msg, _ := res["error_message"].(string)
				return res, "failed", fmt.Errorf("order failed: %s", msg)
			}
			return res, state, nil
		},
		Timeout: 45 * time.Minute,
		Delay: 10 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	rawResult, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Order Failed", err.Error())
		return
	}
	
	result := rawResult.(map[string]interface{})
	marketplaceResourceUUID, ok := result["marketplace_resource_uuid"].(string)
	if !ok || marketplaceResourceUUID == "" {
		resp.Diagnostics.AddError("Resource UUID Missing", "Order completed but marketplace_resource_uuid is missing")
		return
	}

	// Get the marketplace resource to find the actual resource UUID
	var marketplaceRes map[string]interface{}
	err = r.client.GetByUUID(ctx, "/api/marketplace-resources/{uuid}/", marketplaceResourceUUID, &marketplaceRes)
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Marketplace Resource", err.Error())
		return
	}

	// Extract the actual resource UUID from the marketplace resource
	resourceUUID, ok := marketplaceRes["resource_uuid"].(string)
	if !ok || resourceUUID == "" {
		resp.Diagnostics.AddError("Resource UUID Missing", "Marketplace resource does not have resource_uuid")
		return
	}
	data.UUID = types.StringValue(resourceUUID)

	// Store marketplace resource UUID for terminate operations
	data.MarketplaceResourceUuid = types.StringValue(marketplaceResourceUUID)

	// Fetch final resource state
	var finalState map[string]interface{}
	err = r.client.GetByUUID(ctx, "/api/openstack-instances/{uuid}/", resourceUUID, &finalState)
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource", err.Error())
		return
	}
	sourceMap := finalState
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
	if val, ok := sourceMap["availability_zone_name"]; ok {
		if val == nil {
			data.AvailabilityZoneName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.AvailabilityZoneName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["backend_id"]; ok {
		if val == nil {
			data.BackendId = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
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
	if val, ok := sourceMap["customer"]; ok {
		if val == nil {
			data.Customer = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Customer = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["customer_abbreviation"]; ok {
		if val == nil {
			data.CustomerAbbreviation = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerAbbreviation = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["customer_name"]; ok {
		if val == nil {
			data.CustomerName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["customer_native_name"]; ok {
		if val == nil {
			data.CustomerNativeName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerNativeName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["customer_uuid"]; ok {
		if val == nil {
			data.CustomerUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerUuid = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["description"]; ok {
		if val == nil {
			data.Description = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
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
	if val, ok := sourceMap["project"]; ok {
		if val == nil {
			data.Project = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Project = types.StringValue(uuid)
		}
		}
	}
	if val, ok := sourceMap["project_name"]; ok {
		if val == nil {
			data.ProjectName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ProjectName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["project_uuid"]; ok {
		if val == nil {
			data.ProjectUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ProjectUuid = types.StringValue(str)
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
	if val, ok := sourceMap["runtime_state"]; ok {
		if val == nil {
			data.RuntimeState = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.RuntimeState = types.StringValue(str)
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
	if val, ok := sourceMap["service_settings_uuid"]; ok {
		if val == nil {
			data.ServiceSettingsUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ServiceSettingsUuid = types.StringValue(str)
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
	if val, ok := sourceMap["tenant"]; ok {
		if val == nil {
			data.Tenant = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Tenant = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["tenant_uuid"]; ok {
		if val == nil {
			data.TenantUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.TenantUuid = types.StringValue(str)
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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpenstackInstanceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource
	var result map[string]interface{}
	retrievePath := strings.Replace("/api/openstack-instances/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)
	
	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OpenstackInstance",
			"An error occurred while reading the openstack_instance: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
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
	if val, ok := sourceMap["availability_zone_name"]; ok {
		if val == nil {
			data.AvailabilityZoneName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.AvailabilityZoneName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["backend_id"]; ok {
		if val == nil {
			data.BackendId = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
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
	if val, ok := sourceMap["customer"]; ok {
		if val == nil {
			data.Customer = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Customer = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["customer_abbreviation"]; ok {
		if val == nil {
			data.CustomerAbbreviation = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerAbbreviation = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["customer_name"]; ok {
		if val == nil {
			data.CustomerName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["customer_native_name"]; ok {
		if val == nil {
			data.CustomerNativeName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerNativeName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["customer_uuid"]; ok {
		if val == nil {
			data.CustomerUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerUuid = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["description"]; ok {
		if val == nil {
			data.Description = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
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
	if val, ok := sourceMap["project"]; ok {
		if val == nil {
			data.Project = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Project = types.StringValue(uuid)
		}
		}
	}
	if val, ok := sourceMap["project_name"]; ok {
		if val == nil {
			data.ProjectName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ProjectName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["project_uuid"]; ok {
		if val == nil {
			data.ProjectUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ProjectUuid = types.StringValue(str)
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
	if val, ok := sourceMap["runtime_state"]; ok {
		if val == nil {
			data.RuntimeState = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.RuntimeState = types.StringValue(str)
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
	if val, ok := sourceMap["service_settings_uuid"]; ok {
		if val == nil {
			data.ServiceSettingsUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ServiceSettingsUuid = types.StringValue(str)
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
	if val, ok := sourceMap["tenant"]; ok {
		if val == nil {
			data.Tenant = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Tenant = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["tenant_uuid"]; ok {
		if val == nil {
			data.TenantUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.TenantUuid = types.StringValue(str)
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpenstackInstanceResourceModel
	var state OpenstackInstanceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	
	// Read current state to get the UUID (which is computed and not in plan)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}
	
	// Use UUID from state
	data.UUID = state.UUID
	// Phase 1: Standard PATCH (Simple fields)
	patchPayload := map[string]interface{}{}
	if !data.Description.IsNull() && !data.Description.Equal(state.Description) {
		patchPayload["description"] = data.Description.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.Equal(state.Name) {
		patchPayload["name"] = data.Name.ValueString()
	}

	if len(patchPayload) > 0 {
		var result map[string]interface{}
		err := r.client.Update(ctx, "/api/openstack-instances/{uuid}/", data.UUID.ValueString(), patchPayload, &result)
		if err != nil {
			resp.Diagnostics.AddError("Update Failed", err.Error())
			return
		}
	}

	// Phase 2: RPC Actions

	// Fetch updated state
	var result map[string]interface{}
	err := r.client.GetByUUID(ctx, "/api/openstack-instances/{uuid}/", data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Resource After Update", err.Error())
		return
	}
	sourceMap := result
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
	if val, ok := sourceMap["availability_zone_name"]; ok {
		if val == nil {
			data.AvailabilityZoneName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.AvailabilityZoneName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["backend_id"]; ok {
		if val == nil {
			data.BackendId = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
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
	if val, ok := sourceMap["customer"]; ok {
		if val == nil {
			data.Customer = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Customer = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["customer_abbreviation"]; ok {
		if val == nil {
			data.CustomerAbbreviation = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerAbbreviation = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["customer_name"]; ok {
		if val == nil {
			data.CustomerName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["customer_native_name"]; ok {
		if val == nil {
			data.CustomerNativeName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerNativeName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["customer_uuid"]; ok {
		if val == nil {
			data.CustomerUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerUuid = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["description"]; ok {
		if val == nil {
			data.Description = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Description = types.StringValue(str)
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
	if val, ok := sourceMap["project"]; ok {
		if val == nil {
			data.Project = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			// Normalize URL to UUID
			parts := strings.Split(strings.TrimRight(str, "/"), "/")
			uuid := parts[len(parts)-1]
			data.Project = types.StringValue(uuid)
		}
		}
	}
	if val, ok := sourceMap["project_name"]; ok {
		if val == nil {
			data.ProjectName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ProjectName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["project_uuid"]; ok {
		if val == nil {
			data.ProjectUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ProjectUuid = types.StringValue(str)
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
	if val, ok := sourceMap["runtime_state"]; ok {
		if val == nil {
			data.RuntimeState = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.RuntimeState = types.StringValue(str)
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
	if val, ok := sourceMap["service_settings_uuid"]; ok {
		if val == nil {
			data.ServiceSettingsUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.ServiceSettingsUuid = types.StringValue(str)
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
	if val, ok := sourceMap["tenant"]; ok {
		if val == nil {
			data.Tenant = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Tenant = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["tenant_uuid"]; ok {
		if val == nil {
			data.TenantUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.TenantUuid = types.StringValue(str)
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpenstackInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OpenstackInstanceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Order-based Delete
	payload := map[string]interface{}{}

	// Use marketplace_resource_uuid for termination
	marketplaceUUID := data.MarketplaceResourceUuid.ValueString()
	if marketplaceUUID == "" {
		resp.Diagnostics.AddError("Termination Failed", "Marketplace resource UUID is not set")
		return
	}
	url := fmt.Sprintf("/api/marketplace-resources/%s/terminate/", marketplaceUUID)
	var res map[string]interface{}
	err := r.client.Post(ctx, url, payload, &res)
	if err != nil {
		resp.Diagnostics.AddError("Termination Failed", err.Error())
		return
	}
	
	// Wait for deletion if order UUID is returned
	if orderUUID, ok := res["uuid"].(string); ok {
		stateConf := &retry.StateChangeConf{
			Pending: []string{"pending", "executing", "created"},
			Target:  []string{"done"},
			Refresh: func() (interface{}, string, error) {
				var res map[string]interface{}
				err := r.client.GetByUUID(ctx, "/api/marketplace-orders/{uuid}/", orderUUID, &res)
				if err != nil {
					return nil, "", err
				}
				state, _ := res["state"].(string)
				if state == "erred" || state == "rejected" {
					return res, "failed", fmt.Errorf("termination order failed")
				}
				return res, state, nil
			},
			Timeout: 45 * time.Minute,
			Delay: 10 * time.Second,
			MinTimeout: 5 * time.Second,
		}
		_, err := stateConf.WaitForStateContext(ctx)
		if err != nil {
			resp.Diagnostics.AddError("Termination Order Failed", err.Error())
			return
		}
	}
}

func (r *OpenstackInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// TODO: Implement import logic
	resp.Diagnostics.AddError("Not Implemented", "Import is not yet implemented for this resource")
}
