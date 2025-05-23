using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListGroupUsers.json
// this example is just showing the usage of "GroupUser_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementGroupResource created on azure
// for more information of creating ApiManagementGroupResource, please refer to the document of ApiManagementGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string groupId = "57d2ef278aa04f0888cba3f3";
ResourceIdentifier apiManagementGroupResourceId = ApiManagementGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, groupId);
ApiManagementGroupResource apiManagementGroup = client.GetApiManagementGroupResource(apiManagementGroupResourceId);

// invoke the operation and iterate over the result
await foreach (ApiManagementGroupUserData item in apiManagementGroup.GetGroupUsersAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
