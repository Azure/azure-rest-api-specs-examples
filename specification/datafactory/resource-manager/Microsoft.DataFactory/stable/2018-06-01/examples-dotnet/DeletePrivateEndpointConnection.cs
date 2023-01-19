using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataFactory;
using Azure.ResourceManager.DataFactory.Models;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnection_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FactoryPrivateEndpointConnectionResource created on azure
// for more information of creating FactoryPrivateEndpointConnectionResource, please refer to the document of FactoryPrivateEndpointConnectionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string privateEndpointConnectionName = "connection";
ResourceIdentifier factoryPrivateEndpointConnectionResourceId = FactoryPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, privateEndpointConnectionName);
FactoryPrivateEndpointConnectionResource factoryPrivateEndpointConnection = client.GetFactoryPrivateEndpointConnectionResource(factoryPrivateEndpointConnectionResourceId);

// invoke the operation
await factoryPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
