using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPortalSettingsUpdateSignIn.json
// this example is just showing the usage of "SignInSettings_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementPortalSignInSettingResource created on azure
// for more information of creating ApiManagementPortalSignInSettingResource, please refer to the document of ApiManagementPortalSignInSettingResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
ResourceIdentifier apiManagementPortalSignInSettingResourceId = ApiManagementPortalSignInSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
ApiManagementPortalSignInSettingResource apiManagementPortalSignInSetting = client.GetApiManagementPortalSignInSettingResource(apiManagementPortalSignInSettingResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
ApiManagementPortalSignInSettingData data = new ApiManagementPortalSignInSettingData()
{
    IsRedirectEnabled = true,
};
await apiManagementPortalSignInSetting.UpdateAsync(ifMatch, data);

Console.WriteLine($"Succeeded");
