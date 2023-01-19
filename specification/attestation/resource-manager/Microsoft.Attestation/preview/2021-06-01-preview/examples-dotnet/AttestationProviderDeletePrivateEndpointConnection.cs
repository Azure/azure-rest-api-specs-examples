using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Attestation;
using Azure.ResourceManager.Attestation.Models;

// Generated from example definition: specification/attestation/resource-manager/Microsoft.Attestation/preview/2021-06-01-preview/examples/AttestationProviderDeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AttestationPrivateEndpointConnectionResource created on azure
// for more information of creating AttestationPrivateEndpointConnectionResource, please refer to the document of AttestationPrivateEndpointConnectionResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res6977";
string providerName = "sto2527";
string privateEndpointConnectionName = "{privateEndpointConnectionName}";
ResourceIdentifier attestationPrivateEndpointConnectionResourceId = AttestationPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, providerName, privateEndpointConnectionName);
AttestationPrivateEndpointConnectionResource attestationPrivateEndpointConnection = client.GetAttestationPrivateEndpointConnectionResource(attestationPrivateEndpointConnectionResourceId);

// invoke the operation
await attestationPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
