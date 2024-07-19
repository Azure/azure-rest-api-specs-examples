using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteProductTag.json
// this example is just showing the usage of "Tag_DetachFromProduct" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementProductTagResource created on azure
// for more information of creating ApiManagementProductTagResource, please refer to the document of ApiManagementProductTagResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string productId = "59d5b28d1f7fab116c282650";
string tagId = "59d5b28e1f7fab116402044e";
ResourceIdentifier apiManagementProductTagResourceId = ApiManagementProductTagResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, productId, tagId);
ApiManagementProductTagResource apiManagementProductTag = client.GetApiManagementProductTagResource(apiManagementProductTagResourceId);

// invoke the operation
await apiManagementProductTag.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
