using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdateAuthorizationServer.json
// this example is just showing the usage of "AuthorizationServer_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementAuthorizationServerResource created on azure
// for more information of creating ApiManagementAuthorizationServerResource, please refer to the document of ApiManagementAuthorizationServerResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string authsid = "newauthServer";
ResourceIdentifier apiManagementAuthorizationServerResourceId = ApiManagementAuthorizationServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, authsid);
ApiManagementAuthorizationServerResource apiManagementAuthorizationServer = client.GetApiManagementAuthorizationServerResource(apiManagementAuthorizationServerResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
ApiManagementAuthorizationServerPatch patch = new ApiManagementAuthorizationServerPatch
{
    UseInTestConsole = false,
    UseInApiDocumentation = true,
    ClientId = "update",
    ClientSecret = "updated",
};
ApiManagementAuthorizationServerResource result = await apiManagementAuthorizationServer.UpdateAsync(ifMatch, patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiManagementAuthorizationServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
