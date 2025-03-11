using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementListUserIdentities.json
// this example is just showing the usage of "UserIdentities_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementUserResource created on azure
// for more information of creating ApiManagementUserResource, please refer to the document of ApiManagementUserResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string userId = "57f2af53bb17172280f44057";
ResourceIdentifier apiManagementUserResourceId = ApiManagementUserResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, userId);
ApiManagementUserResource apiManagementUser = client.GetApiManagementUserResource(apiManagementUserResourceId);

// invoke the operation and iterate over the result
await foreach (UserIdentityContract item in apiManagementUser.GetUserIdentitiesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
