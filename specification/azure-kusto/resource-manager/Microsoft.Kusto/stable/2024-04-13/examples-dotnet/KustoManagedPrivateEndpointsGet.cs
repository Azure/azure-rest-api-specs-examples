using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Kusto;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoManagedPrivateEndpointsGet.json
// this example is just showing the usage of "ManagedPrivateEndpoints_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoManagedPrivateEndpointResource created on azure
// for more information of creating KustoManagedPrivateEndpointResource, please refer to the document of KustoManagedPrivateEndpointResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster";
string managedPrivateEndpointName = "managedPrivateEndpointTest";
ResourceIdentifier kustoManagedPrivateEndpointResourceId = KustoManagedPrivateEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, managedPrivateEndpointName);
KustoManagedPrivateEndpointResource kustoManagedPrivateEndpoint = client.GetKustoManagedPrivateEndpointResource(kustoManagedPrivateEndpointResourceId);

// invoke the operation
KustoManagedPrivateEndpointResource result = await kustoManagedPrivateEndpoint.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KustoManagedPrivateEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
