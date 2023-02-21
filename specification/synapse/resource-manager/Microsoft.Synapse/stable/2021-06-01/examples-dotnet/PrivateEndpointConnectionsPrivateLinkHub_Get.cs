using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/PrivateEndpointConnectionsPrivateLinkHub_Get.json
// this example is just showing the usage of "PrivateEndpointConnectionsPrivateLinkHub_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapsePrivateLinkHubResource created on azure
// for more information of creating SynapsePrivateLinkHubResource, please refer to the document of SynapsePrivateLinkHubResource
string subscriptionId = "48b08652-d7a1-4d52-b13f-5a2471dce57b";
string resourceGroupName = "gh-res-grp";
string privateLinkHubName = "pe0";
ResourceIdentifier synapsePrivateLinkHubResourceId = SynapsePrivateLinkHubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateLinkHubName);
SynapsePrivateLinkHubResource synapsePrivateLinkHub = client.GetSynapsePrivateLinkHubResource(synapsePrivateLinkHubResourceId);

// get the collection of this SynapsePrivateEndpointConnectionForPrivateLinkHubResource
SynapsePrivateEndpointConnectionForPrivateLinkHubCollection collection = synapsePrivateLinkHub.GetSynapsePrivateEndpointConnectionForPrivateLinkHubs();

// invoke the operation
string privateEndpointConnectionName = "pe0-f3ed30f5-338c-4855-a542-24a403694ad2";
bool result = await collection.ExistsAsync(privateEndpointConnectionName);

Console.WriteLine($"Succeeded: {result}");
