using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiOperationTag.json
// this example is just showing the usage of "Tag_AssignToOperation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiOperationResource created on azure
// for more information of creating ApiOperationResource, please refer to the document of ApiOperationResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "5931a75ae4bbd512a88c680b";
string operationId = "5931a75ae4bbd512a88c680a";
ResourceIdentifier apiOperationResourceId = ApiOperationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, operationId);
ApiOperationResource apiOperation = client.GetApiOperationResource(apiOperationResourceId);

// get the collection of this ApiOperationTagResource
ApiOperationTagCollection collection = apiOperation.GetApiOperationTags();

// invoke the operation
string tagId = "tagId1";
ArmOperation<ApiOperationTagResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, tagId);
ApiOperationTagResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TagContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
