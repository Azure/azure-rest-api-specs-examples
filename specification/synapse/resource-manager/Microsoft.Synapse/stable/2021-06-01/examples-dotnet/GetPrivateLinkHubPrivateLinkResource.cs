using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetPrivateLinkHubPrivateLinkResource.json
// this example is just showing the usage of "PrivateLinkHubPrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapsePrivateLinkHubResource created on azure
// for more information of creating SynapsePrivateLinkHubResource, please refer to the document of SynapsePrivateLinkHubResource
string subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
string resourceGroupName = "ExampleResourceGroup";
string privateLinkHubName = "ExamplePrivateLinkHub";
ResourceIdentifier synapsePrivateLinkHubResourceId = SynapsePrivateLinkHubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateLinkHubName);
SynapsePrivateLinkHubResource synapsePrivateLinkHub = client.GetSynapsePrivateLinkHubResource(synapsePrivateLinkHubResourceId);

// get the collection of this SynapsePrivateLinkResource
SynapsePrivateLinkResourceCollection collection = synapsePrivateLinkHub.GetSynapsePrivateLinkResources();

// invoke the operation
string privateLinkResourceName = "sql";
bool result = await collection.ExistsAsync(privateLinkResourceName);

Console.WriteLine($"Succeeded: {result}");
