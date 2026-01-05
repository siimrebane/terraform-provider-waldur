package resources

import (
	"context"
    "strings"
    

	
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/waldur/terraform-waldur-provider/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &StructureProjectResource{}
var _ resource.ResourceWithImportState = &StructureProjectResource{}

func NewStructureProjectResource() resource.Resource {
	return &StructureProjectResource{}
}

// StructureProjectResource defines the resource implementation.
type StructureProjectResource struct {
	client *client.Client
}

// StructureProjectResourceModel describes the resource data model.
type StructureProjectResourceModel struct {
	UUID types.String `tfsdk:"id"`
	BackendId types.String `tfsdk:"backend_id"`
	Customer types.String `tfsdk:"customer"`
	Description types.String `tfsdk:"description"`
	EndDate types.String `tfsdk:"end_date"`
	GracePeriodDays types.Int64 `tfsdk:"grace_period_days"`
	Image types.String `tfsdk:"image"`
	IsIndustry types.Bool `tfsdk:"is_industry"`
	Name types.String `tfsdk:"name"`
	Slug types.String `tfsdk:"slug"`
	StaffNotes types.String `tfsdk:"staff_notes"`
	StartDate types.String `tfsdk:"start_date"`
	Type types.String `tfsdk:"type"`
	Created types.String `tfsdk:"created"`
	CustomerAbbreviation types.String `tfsdk:"customer_abbreviation"`
	CustomerDisplayBillingInfoInProjects types.Bool `tfsdk:"customer_display_billing_info_in_projects"`
	CustomerName types.String `tfsdk:"customer_name"`
	CustomerNativeName types.String `tfsdk:"customer_native_name"`
	CustomerSlug types.String `tfsdk:"customer_slug"`
	CustomerUuid types.String `tfsdk:"customer_uuid"`
	EndDateRequestedBy types.String `tfsdk:"end_date_requested_by"`
	IsRemoved types.Bool `tfsdk:"is_removed"`
	MaxServiceAccounts types.Int64 `tfsdk:"max_service_accounts"`
	OecdFos2007Label types.String `tfsdk:"oecd_fos_2007_label"`
	ProjectCredit types.Float64 `tfsdk:"project_credit"`
	ResourcesCount types.Int64 `tfsdk:"resources_count"`
	TypeName types.String `tfsdk:"type_name"`
	TypeUuid types.String `tfsdk:"type_uuid"`
	Url types.String `tfsdk:"url"`
}

func (r *StructureProjectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_structure_project"
}

func (r *StructureProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "StructureProject resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Resource UUID (used as Terraform ID)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"backend_id":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer":schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "",
			},
			"description":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project description (HTML content will be sanitized)",
			},
			"end_date":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project end date. Setting this field requires DELETE_PROJECT permission.",
			},
			"grace_period_days":schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Number of extra days after project end date before resources are terminated. Overrides customer-level setting.",
			},
			"image":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"is_industry":schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"name":schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "",
			},
			"slug":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "URL-friendly identifier. Only editable by staff users.",
			},
			"staff_notes":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Internal notes visible only to staff and support users (HTML content will be sanitized)",
			},
			"start_date":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Project start date. Cannot be edited after the start date has arrived.",
			},
			"type":schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "",
			},
			"created":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_abbreviation":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_display_billing_info_in_projects":schema.BoolAttribute{
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
			"customer_slug":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"customer_uuid":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"end_date_requested_by":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"is_removed":schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"max_service_accounts":schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Maximum number of service accounts allowed",
			},
			"oecd_fos_2007_label":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Human-readable label for the OECD FOS 2007 classification code",
			},
			"project_credit":schema.Float64Attribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"resources_count":schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Number of active resources in this project",
			},
			"type_name":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"type_uuid":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
			"url":schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "",
			},
		},
	}
}

func (r *StructureProjectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *StructureProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StructureProjectResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Prepare request body
	requestBody := map[string]interface{}{
		"backend_id": data.BackendId.ValueString(),
		"customer": data.Customer.ValueString(),
		"description": data.Description.ValueString(),
		"end_date": data.EndDate.ValueString(),
		"grace_period_days": data.GracePeriodDays.ValueInt64(),
		"image": data.Image.ValueString(),
		"is_industry": data.IsIndustry.ValueBool(),
		"name": data.Name.ValueString(),
		"slug": data.Slug.ValueString(),
		"staff_notes": data.StaffNotes.ValueString(),
		"start_date": data.StartDate.ValueString(),
		"type": data.Type.ValueString(),
	}

	// Call Waldur API to create resource
	var result map[string]interface{}
	
	err := r.client.Create(ctx, "/api/projects/", requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create StructureProject",
			"An error occurred while creating the structure_project: "+err.Error(),
		)
		return
	}

	// Extract UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	if val, ok := sourceMap["backend_id"]; ok {
		if val == nil {
			data.BackendId = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
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
	if val, ok := sourceMap["customer_display_billing_info_in_projects"]; ok {
		if val == nil {
			data.CustomerDisplayBillingInfoInProjects = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.CustomerDisplayBillingInfoInProjects = types.BoolValue(b)
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
	if val, ok := sourceMap["customer_slug"]; ok {
		if val == nil {
			data.CustomerSlug = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerSlug = types.StringValue(str)
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
	if val, ok := sourceMap["end_date"]; ok {
		if val == nil {
			data.EndDate = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.EndDate = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["end_date_requested_by"]; ok {
		if val == nil {
			data.EndDateRequestedBy = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.EndDateRequestedBy = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["grace_period_days"]; ok {
		if val == nil {
			data.GracePeriodDays = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.GracePeriodDays = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["image"]; ok {
		if val == nil {
			data.Image = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["is_industry"]; ok {
		if val == nil {
			data.IsIndustry = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.IsIndustry = types.BoolValue(b)
		}
		}
	}
	if val, ok := sourceMap["is_removed"]; ok {
		if val == nil {
			data.IsRemoved = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.IsRemoved = types.BoolValue(b)
		}
		}
	}
	if val, ok := sourceMap["max_service_accounts"]; ok {
		if val == nil {
			data.MaxServiceAccounts = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.MaxServiceAccounts = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["oecd_fos_2007_label"]; ok {
		if val == nil {
			data.OecdFos2007Label = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.OecdFos2007Label = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["project_credit"]; ok {
		if val == nil {
			data.ProjectCredit = types.Float64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.ProjectCredit = types.Float64Value(num)
		}
		}
	}
	if val, ok := sourceMap["resources_count"]; ok {
		if val == nil {
			data.ResourcesCount = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.ResourcesCount = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["slug"]; ok {
		if val == nil {
			data.Slug = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["staff_notes"]; ok {
		if val == nil {
			data.StaffNotes = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.StaffNotes = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["start_date"]; ok {
		if val == nil {
			data.StartDate = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.StartDate = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["type"]; ok {
		if val == nil {
			data.Type = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Type = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["type_name"]; ok {
		if val == nil {
			data.TypeName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.TypeName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["type_uuid"]; ok {
		if val == nil {
			data.TypeUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.TypeUuid = types.StringValue(str)
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

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data StructureProjectResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call Waldur API to read resource
	var result map[string]interface{}
	retrievePath := strings.Replace("/api/projects/{uuid}/", "{uuid}", data.UUID.ValueString(), 1)
	
	err := r.client.GetByUUID(ctx, retrievePath, data.UUID.ValueString(), &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read StructureProject",
			"An error occurred while reading the structure_project: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	if val, ok := sourceMap["backend_id"]; ok {
		if val == nil {
			data.BackendId = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
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
	if val, ok := sourceMap["customer_display_billing_info_in_projects"]; ok {
		if val == nil {
			data.CustomerDisplayBillingInfoInProjects = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.CustomerDisplayBillingInfoInProjects = types.BoolValue(b)
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
	if val, ok := sourceMap["customer_slug"]; ok {
		if val == nil {
			data.CustomerSlug = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerSlug = types.StringValue(str)
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
	if val, ok := sourceMap["end_date"]; ok {
		if val == nil {
			data.EndDate = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.EndDate = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["end_date_requested_by"]; ok {
		if val == nil {
			data.EndDateRequestedBy = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.EndDateRequestedBy = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["grace_period_days"]; ok {
		if val == nil {
			data.GracePeriodDays = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.GracePeriodDays = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["image"]; ok {
		if val == nil {
			data.Image = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["is_industry"]; ok {
		if val == nil {
			data.IsIndustry = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.IsIndustry = types.BoolValue(b)
		}
		}
	}
	if val, ok := sourceMap["is_removed"]; ok {
		if val == nil {
			data.IsRemoved = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.IsRemoved = types.BoolValue(b)
		}
		}
	}
	if val, ok := sourceMap["max_service_accounts"]; ok {
		if val == nil {
			data.MaxServiceAccounts = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.MaxServiceAccounts = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["oecd_fos_2007_label"]; ok {
		if val == nil {
			data.OecdFos2007Label = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.OecdFos2007Label = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["project_credit"]; ok {
		if val == nil {
			data.ProjectCredit = types.Float64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.ProjectCredit = types.Float64Value(num)
		}
		}
	}
	if val, ok := sourceMap["resources_count"]; ok {
		if val == nil {
			data.ResourcesCount = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.ResourcesCount = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["slug"]; ok {
		if val == nil {
			data.Slug = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["staff_notes"]; ok {
		if val == nil {
			data.StaffNotes = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.StaffNotes = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["start_date"]; ok {
		if val == nil {
			data.StartDate = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.StartDate = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["type"]; ok {
		if val == nil {
			data.Type = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Type = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["type_name"]; ok {
		if val == nil {
			data.TypeName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.TypeName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["type_uuid"]; ok {
		if val == nil {
			data.TypeUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.TypeUuid = types.StringValue(str)
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data StructureProjectResourceModel
	var state StructureProjectResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	
	// Read current state to get the UUID (which is computed and not in plan)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}
	
	// Use UUID from state
	data.UUID = state.UUID
	// Prepare request body
	requestBody := map[string]interface{}{
		"backend_id": data.BackendId.ValueString(),
		"customer": data.Customer.ValueString(),
		"description": data.Description.ValueString(),
		"end_date": data.EndDate.ValueString(),
		"grace_period_days": data.GracePeriodDays.ValueInt64(),
		"image": data.Image.ValueString(),
		"is_industry": data.IsIndustry.ValueBool(),
		"name": data.Name.ValueString(),
		"slug": data.Slug.ValueString(),
		"staff_notes": data.StaffNotes.ValueString(),
		"start_date": data.StartDate.ValueString(),
		"type": data.Type.ValueString(),
	}

	// Call Waldur API to update resource
	var result map[string]interface{}
	
	err := r.client.Update(ctx, "/api/projects/{uuid}/", data.UUID.ValueString(), requestBody, &result)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update StructureProject",
			"An error occurred while updating the structure_project: "+err.Error(),
		)
		return
	}

	// Update UUID from response
	if uuid, ok := result["uuid"].(string); ok {
		data.UUID = types.StringValue(uuid)
	}

	sourceMap := result
	// Map response fields to data model
	if val, ok := sourceMap["backend_id"]; ok {
		if val == nil {
			data.BackendId = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.BackendId = types.StringValue(str)
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
	if val, ok := sourceMap["customer_display_billing_info_in_projects"]; ok {
		if val == nil {
			data.CustomerDisplayBillingInfoInProjects = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.CustomerDisplayBillingInfoInProjects = types.BoolValue(b)
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
	if val, ok := sourceMap["customer_slug"]; ok {
		if val == nil {
			data.CustomerSlug = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.CustomerSlug = types.StringValue(str)
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
	if val, ok := sourceMap["end_date"]; ok {
		if val == nil {
			data.EndDate = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.EndDate = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["end_date_requested_by"]; ok {
		if val == nil {
			data.EndDateRequestedBy = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.EndDateRequestedBy = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["grace_period_days"]; ok {
		if val == nil {
			data.GracePeriodDays = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.GracePeriodDays = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["image"]; ok {
		if val == nil {
			data.Image = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Image = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["is_industry"]; ok {
		if val == nil {
			data.IsIndustry = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.IsIndustry = types.BoolValue(b)
		}
		}
	}
	if val, ok := sourceMap["is_removed"]; ok {
		if val == nil {
			data.IsRemoved = types.BoolNull()
		} else {
		if b, ok := val.(bool); ok {
			data.IsRemoved = types.BoolValue(b)
		}
		}
	}
	if val, ok := sourceMap["max_service_accounts"]; ok {
		if val == nil {
			data.MaxServiceAccounts = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.MaxServiceAccounts = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["oecd_fos_2007_label"]; ok {
		if val == nil {
			data.OecdFos2007Label = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.OecdFos2007Label = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["project_credit"]; ok {
		if val == nil {
			data.ProjectCredit = types.Float64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.ProjectCredit = types.Float64Value(num)
		}
		}
	}
	if val, ok := sourceMap["resources_count"]; ok {
		if val == nil {
			data.ResourcesCount = types.Int64Null()
		} else {
		if num, ok := val.(float64); ok {
			data.ResourcesCount = types.Int64Value(int64(num))
		}
		}
	}
	if val, ok := sourceMap["slug"]; ok {
		if val == nil {
			data.Slug = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Slug = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["staff_notes"]; ok {
		if val == nil {
			data.StaffNotes = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.StaffNotes = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["start_date"]; ok {
		if val == nil {
			data.StartDate = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.StartDate = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["type"]; ok {
		if val == nil {
			data.Type = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.Type = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["type_name"]; ok {
		if val == nil {
			data.TypeName = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.TypeName = types.StringValue(str)
		}
		}
	}
	if val, ok := sourceMap["type_uuid"]; ok {
		if val == nil {
			data.TypeUuid = types.StringNull()
		} else {
		if str, ok := val.(string); ok {
			data.TypeUuid = types.StringValue(str)
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

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StructureProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data StructureProjectResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// Call Waldur API to delete resource
	err := r.client.DeleteByUUID(ctx, "/api/projects/{uuid}/", data.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete StructureProject",
			"An error occurred while deleting the structure_project: "+err.Error(),
		)
		return
	}
}

func (r *StructureProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// TODO: Implement import logic
	resp.Diagnostics.AddError("Not Implemented", "Import is not yet implemented for this resource")
}
