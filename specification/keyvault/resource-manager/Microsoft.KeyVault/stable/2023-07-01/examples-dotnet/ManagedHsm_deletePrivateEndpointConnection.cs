using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.KeyVault.Models;
using Azure.ResourceManager.KeyVault;

// Generated from example definition: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/ManagedHsm_deletePrivateEndpointConnection.json
// this example is just showing the usage of "MHSMPrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedHsmPrivateEndpointConnectionResource created on azure
// for more information of creating ManagedHsmPrivateEndpointConnectionResource, please refer to the document of ManagedHsmPrivateEndpointConnectionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "sample-group";
string name = "sample-mhsm";
string privateEndpointConnectionName = "sample-pec";
ResourceIdentifier managedHsmPrivateEndpointConnectionResourceId = ManagedHsmPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, privateEndpointConnectionName);
ManagedHsmPrivateEndpointConnectionResource managedHsmPrivateEndpointConnection = client.GetManagedHsmPrivateEndpointConnectionResource(managedHsmPrivateEndpointConnectionResourceId);

// invoke the operation
ArmOperation<ManagedHsmPrivateEndpointConnectionResource> lro = await managedHsmPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);
ManagedHsmPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedHsmPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
