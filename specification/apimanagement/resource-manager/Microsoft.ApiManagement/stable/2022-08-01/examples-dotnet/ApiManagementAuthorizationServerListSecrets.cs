using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementAuthorizationServerListSecrets.json
// this example is just showing the usage of "AuthorizationServer_ListSecrets" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementAuthorizationServerResource created on azure
// for more information of creating ApiManagementAuthorizationServerResource, please refer to the document of ApiManagementAuthorizationServerResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string authsid = "newauthServer2";
ResourceIdentifier apiManagementAuthorizationServerResourceId = ApiManagementAuthorizationServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, authsid);
ApiManagementAuthorizationServerResource apiManagementAuthorizationServer = client.GetApiManagementAuthorizationServerResource(apiManagementAuthorizationServerResourceId);

// invoke the operation
AuthorizationServerSecretsContract result = await apiManagementAuthorizationServer.GetSecretsAsync();

Console.WriteLine($"Succeeded: {result}");
