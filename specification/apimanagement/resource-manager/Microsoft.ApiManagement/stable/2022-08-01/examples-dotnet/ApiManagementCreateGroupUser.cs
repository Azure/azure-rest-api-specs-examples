using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGroupUser.json
// this example is just showing the usage of "GroupUser_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementGroupResource created on azure
// for more information of creating ApiManagementGroupResource, please refer to the document of ApiManagementGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string groupId = "tempgroup";
ResourceIdentifier apiManagementGroupResourceId = ApiManagementGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, groupId);
ApiManagementGroupResource apiManagementGroup = client.GetApiManagementGroupResource(apiManagementGroupResourceId);

// invoke the operation
string userId = "59307d350af58404d8a26300";
ApiManagementGroupUserData result = await apiManagementGroup.CreateGroupUserAsync(userId);

Console.WriteLine($"Succeeded: {result}");
