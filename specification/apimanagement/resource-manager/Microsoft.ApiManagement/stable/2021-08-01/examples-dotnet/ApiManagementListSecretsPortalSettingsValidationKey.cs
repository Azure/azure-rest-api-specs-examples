using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApiManagement;
using Azure.ResourceManager.ApiManagement.Models;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListSecretsPortalSettingsValidationKey.json
// this example is just showing the usage of "DelegationSettings_ListSecrets" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementPortalDelegationSettingResource created on azure
// for more information of creating ApiManagementPortalDelegationSettingResource, please refer to the document of ApiManagementPortalDelegationSettingResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
ResourceIdentifier apiManagementPortalDelegationSettingResourceId = ApiManagementPortalDelegationSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
ApiManagementPortalDelegationSettingResource apiManagementPortalDelegationSetting = client.GetApiManagementPortalDelegationSettingResource(apiManagementPortalDelegationSettingResourceId);

// invoke the operation
PortalSettingValidationKeyContract result = await apiManagementPortalDelegationSetting.GetSecretsAsync();

Console.WriteLine($"Succeeded: {result}");
