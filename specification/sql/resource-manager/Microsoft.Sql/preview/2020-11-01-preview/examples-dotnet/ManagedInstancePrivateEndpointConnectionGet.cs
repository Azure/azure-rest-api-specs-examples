using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstancePrivateEndpointConnectionGet.json
// this example is just showing the usage of "ManagedInstancePrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstancePrivateEndpointConnectionResource created on azure
// for more information of creating ManagedInstancePrivateEndpointConnectionResource, please refer to the document of ManagedInstancePrivateEndpointConnectionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string managedInstanceName = "test-cl";
string privateEndpointConnectionName = "private-endpoint-connection-name";
ResourceIdentifier managedInstancePrivateEndpointConnectionResourceId = ManagedInstancePrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, privateEndpointConnectionName);
ManagedInstancePrivateEndpointConnectionResource managedInstancePrivateEndpointConnection = client.GetManagedInstancePrivateEndpointConnectionResource(managedInstancePrivateEndpointConnectionResourceId);

// invoke the operation
ManagedInstancePrivateEndpointConnectionResource result = await managedInstancePrivateEndpointConnection.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedInstancePrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
