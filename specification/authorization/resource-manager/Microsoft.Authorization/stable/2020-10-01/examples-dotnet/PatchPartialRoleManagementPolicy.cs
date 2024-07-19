using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Authorization.Models;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/PatchPartialRoleManagementPolicy.json
// this example is just showing the usage of "RoleManagementPolicies_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RoleManagementPolicyResource created on azure
// for more information of creating RoleManagementPolicyResource, please refer to the document of RoleManagementPolicyResource
string scope = "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368";
string roleManagementPolicyName = "570c3619-7688-4b34-b290-2b8bb3ccab2a";
ResourceIdentifier roleManagementPolicyResourceId = RoleManagementPolicyResource.CreateResourceIdentifier(scope, roleManagementPolicyName);
RoleManagementPolicyResource roleManagementPolicy = client.GetRoleManagementPolicyResource(roleManagementPolicyResourceId);

// invoke the operation
RoleManagementPolicyData data = new RoleManagementPolicyData()
{
    Rules =
    {
    new RoleManagementPolicyExpirationRule()
    {
    IsExpirationRequired = false,
    MaximumDuration = XmlConvert.ToTimeSpan("P180D"),
    Id = "Expiration_Admin_Eligibility",
    Target = new RoleManagementPolicyRuleTarget()
    {
    Caller = "Admin",
    Operations =
    {
    "All"
    },
    Level = RoleManagementAssignmentLevel.Eligibility,
    TargetObjects =
    {
    },
    InheritableSettings =
    {
    },
    EnforcedSettings =
    {
    },
    },
    },new RoleManagementPolicyNotificationRule()
    {
    NotificationDeliveryType = NotificationDeliveryType.Email,
    NotificationLevel = RoleManagementPolicyNotificationLevel.Critical,
    RecipientType = RoleManagementPolicyRecipientType.Admin,
    NotificationRecipients =
    {
    "admin_admin_eligible@test.com"
    },
    AreDefaultRecipientsEnabled = false,
    Id = "Notification_Admin_Admin_Eligibility",
    Target = new RoleManagementPolicyRuleTarget()
    {
    Caller = "Admin",
    Operations =
    {
    "All"
    },
    Level = RoleManagementAssignmentLevel.Eligibility,
    TargetObjects =
    {
    },
    InheritableSettings =
    {
    },
    EnforcedSettings =
    {
    },
    },
    }
    },
};
RoleManagementPolicyResource result = await roleManagementPolicy.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RoleManagementPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
