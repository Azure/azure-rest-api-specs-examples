using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementPortalSettingsUpdateDelegation.json
// this example is just showing the usage of "DelegationSettings_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementPortalDelegationSettingResource created on azure
// for more information of creating ApiManagementPortalDelegationSettingResource, please refer to the document of ApiManagementPortalDelegationSettingResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
ResourceIdentifier apiManagementPortalDelegationSettingResourceId = ApiManagementPortalDelegationSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
ApiManagementPortalDelegationSettingResource apiManagementPortalDelegationSetting = client.GetApiManagementPortalDelegationSettingResource(apiManagementPortalDelegationSettingResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
ApiManagementPortalDelegationSettingData data = new ApiManagementPortalDelegationSettingData
{
    Uri = new Uri("http://contoso.com/delegation"),
    ValidationKey = "<validationKey>",
    IsSubscriptionDelegationEnabled = true,
    IsUserRegistrationDelegationEnabled = true,
};
await apiManagementPortalDelegationSetting.UpdateAsync(ifMatch, data);

Console.WriteLine("Succeeded");
