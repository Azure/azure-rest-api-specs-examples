using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PrivateEndpointConnectionGet.json
// this example is just showing the usage of "PrivateEndpointConnection_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchAccountResource created on azure
// for more information of creating BatchAccountResource, please refer to the document of BatchAccountResource
string subscriptionId = "subid";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
ResourceIdentifier batchAccountResourceId = BatchAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
BatchAccountResource batchAccount = client.GetBatchAccountResource(batchAccountResourceId);

// get the collection of this BatchPrivateEndpointConnectionResource
BatchPrivateEndpointConnectionCollection collection = batchAccount.GetBatchPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "testprivateEndpointConnection5testprivateEndpointConnection5.24d6b4b5-e65c-4330-bbe9-3a290d62f8e0";
bool result = await collection.ExistsAsync(privateEndpointConnectionName);

Console.WriteLine($"Succeeded: {result}");
