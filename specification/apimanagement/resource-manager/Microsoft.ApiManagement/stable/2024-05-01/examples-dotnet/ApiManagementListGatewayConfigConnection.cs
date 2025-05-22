using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListGatewayConfigConnection.json
// this example is just showing the usage of "ApiGatewayConfigConnection_ListByGateway" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiGatewayResource created on azure
// for more information of creating ApiGatewayResource, please refer to the document of ApiGatewayResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string gatewayName = "standard-gw-1";
ResourceIdentifier apiGatewayResourceId = ApiGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, gatewayName);
ApiGatewayResource apiGateway = client.GetApiGatewayResource(apiGatewayResourceId);

// get the collection of this ApiGatewayConfigConnectionResource
ApiGatewayConfigConnectionCollection collection = apiGateway.GetApiGatewayConfigConnections();

// invoke the operation and iterate over the result
await foreach (ApiGatewayConfigConnectionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ApiGatewayConfigConnectionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
