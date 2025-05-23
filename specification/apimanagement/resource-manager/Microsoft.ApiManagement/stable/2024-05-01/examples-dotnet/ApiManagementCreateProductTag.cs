using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateProductTag.json
// this example is just showing the usage of "Tag_AssignToProduct" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementProductTagResource created on azure
// for more information of creating ApiManagementProductTagResource, please refer to the document of ApiManagementProductTagResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string productId = "5931a75ae4bbd512a88c680b";
string tagId = "tagId1";
ResourceIdentifier apiManagementProductTagResourceId = ApiManagementProductTagResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, productId, tagId);
ApiManagementProductTagResource apiManagementProductTag = client.GetApiManagementProductTagResource(apiManagementProductTagResourceId);

// invoke the operation
ArmOperation<ApiManagementProductTagResource> lro = await apiManagementProductTag.UpdateAsync(WaitUntil.Completed);
ApiManagementProductTagResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TagContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
