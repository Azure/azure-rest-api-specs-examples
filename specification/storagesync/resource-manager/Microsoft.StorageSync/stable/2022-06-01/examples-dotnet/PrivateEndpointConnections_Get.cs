using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.StorageSync;
using Azure.ResourceManager.StorageSync.Models;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/PrivateEndpointConnections_Get.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageSyncPrivateEndpointConnectionResource created on azure
// for more information of creating StorageSyncPrivateEndpointConnectionResource, please refer to the document of StorageSyncPrivateEndpointConnectionResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res6977";
string storageSyncServiceName = "sss2527";
string privateEndpointConnectionName = "{privateEndpointConnectionName}";
ResourceIdentifier storageSyncPrivateEndpointConnectionResourceId = StorageSyncPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageSyncServiceName, privateEndpointConnectionName);
StorageSyncPrivateEndpointConnectionResource storageSyncPrivateEndpointConnection = client.GetStorageSyncPrivateEndpointConnectionResource(storageSyncPrivateEndpointConnectionResourceId);

// invoke the operation
StorageSyncPrivateEndpointConnectionResource result = await storageSyncPrivateEndpointConnection.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageSyncPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
