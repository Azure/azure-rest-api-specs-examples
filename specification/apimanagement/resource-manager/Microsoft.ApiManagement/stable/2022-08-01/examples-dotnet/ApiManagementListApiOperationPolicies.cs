using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListApiOperationPolicies.json
// this example is just showing the usage of "ApiOperationPolicy_ListByOperation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiOperationResource created on azure
// for more information of creating ApiOperationResource, please refer to the document of ApiOperationResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "599e2953193c3c0bd0b3e2fa";
string operationId = "599e29ab193c3c0bd0b3e2fb";
ResourceIdentifier apiOperationResourceId = ApiOperationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, operationId);
ApiOperationResource apiOperation = client.GetApiOperationResource(apiOperationResourceId);

// get the collection of this ApiOperationPolicyResource
ApiOperationPolicyCollection collection = apiOperation.GetApiOperationPolicies();

// invoke the operation and iterate over the result
await foreach (ApiOperationPolicyResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PolicyContractData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
