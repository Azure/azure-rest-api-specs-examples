using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPortalSettingsUpdateSignUp.json
// this example is just showing the usage of "SignUpSettings_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementPortalSignUpSettingResource created on azure
// for more information of creating ApiManagementPortalSignUpSettingResource, please refer to the document of ApiManagementPortalSignUpSettingResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
ResourceIdentifier apiManagementPortalSignUpSettingResourceId = ApiManagementPortalSignUpSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
ApiManagementPortalSignUpSettingResource apiManagementPortalSignUpSetting = client.GetApiManagementPortalSignUpSettingResource(apiManagementPortalSignUpSettingResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
ApiManagementPortalSignUpSettingData data = new ApiManagementPortalSignUpSettingData()
{
    IsSignUpDeveloperPortalEnabled = true,
    TermsOfService = new TermsOfServiceProperties()
    {
        Text = "Terms of service text.",
        IsDisplayEnabled = true,
        IsConsentRequired = true,
    },
};
await apiManagementPortalSignUpSetting.UpdateAsync(ifMatch, data);

Console.WriteLine($"Succeeded");
