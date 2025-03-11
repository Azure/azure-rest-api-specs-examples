using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementGetWorkspaceApiOperation.json
// this example is just showing the usage of "WorkspaceApiOperation_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceApiResource created on azure
// for more information of creating ServiceWorkspaceApiResource, please refer to the document of ServiceWorkspaceApiResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
string apiId = "57d2ef278aa04f0888cba3f3";
ResourceIdentifier serviceWorkspaceApiResourceId = ServiceWorkspaceApiResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, apiId);
ServiceWorkspaceApiResource serviceWorkspaceApi = client.GetServiceWorkspaceApiResource(serviceWorkspaceApiResourceId);

// get the collection of this ServiceWorkspaceApiOperationResource
ServiceWorkspaceApiOperationCollection collection = serviceWorkspaceApi.GetServiceWorkspaceApiOperations();

// invoke the operation
string operationId = "57d2ef278aa04f0ad01d6cdc";
NullableResponse<ServiceWorkspaceApiOperationResource> response = await collection.GetIfExistsAsync(operationId);
ServiceWorkspaceApiOperationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ApiOperationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
