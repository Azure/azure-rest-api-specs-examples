using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadTag.json
// this example is just showing the usage of "Tag_GetEntityState" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementTagResource created on azure
// for more information of creating ApiManagementTagResource, please refer to the document of ApiManagementTagResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string tagId = "59306a29e4bbd510dc24e5f9";
ResourceIdentifier apiManagementTagResourceId = ApiManagementTagResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, tagId);
ApiManagementTagResource apiManagementTag = client.GetApiManagementTagResource(apiManagementTagResourceId);

// invoke the operation
bool result = await apiManagementTag.GetEntityStateAsync();

Console.WriteLine($"Succeeded: {result}");
