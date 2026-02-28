using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Batch.Models;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/Batch/stable/2025-06-01/examples/PrivateEndpointConnectionDelete.json
// this example is just showing the usage of "PrivateEndpointConnection_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchPrivateEndpointConnectionResource created on azure
// for more information of creating BatchPrivateEndpointConnectionResource, please refer to the document of BatchPrivateEndpointConnectionResource
string subscriptionId = "12345678-1234-1234-1234-123456789012";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
string privateEndpointConnectionName = "testprivateEndpointConnection5testprivateEndpointConnection5.24d6b4b5-e65c-4330-bbe9-3a290d62f8e0";
ResourceIdentifier batchPrivateEndpointConnectionResourceId = BatchPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, privateEndpointConnectionName);
BatchPrivateEndpointConnectionResource batchPrivateEndpointConnection = client.GetBatchPrivateEndpointConnectionResource(batchPrivateEndpointConnectionResourceId);

// invoke the operation
await batchPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
