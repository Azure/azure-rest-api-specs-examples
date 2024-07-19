using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementIdentityProviderListSecrets.json
// this example is just showing the usage of "IdentityProvider_ListSecrets" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementIdentityProviderResource created on azure
// for more information of creating ApiManagementIdentityProviderResource, please refer to the document of ApiManagementIdentityProviderResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
IdentityProviderType identityProviderName = IdentityProviderType.AadB2C;
ResourceIdentifier apiManagementIdentityProviderResourceId = ApiManagementIdentityProviderResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, identityProviderName);
ApiManagementIdentityProviderResource apiManagementIdentityProvider = client.GetApiManagementIdentityProviderResource(apiManagementIdentityProviderResourceId);

// invoke the operation
ClientSecretContract result = await apiManagementIdentityProvider.GetSecretsAsync();

Console.WriteLine($"Succeeded: {result}");
