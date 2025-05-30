using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetPrivateLinkHubPrivateLinkResource.json
// this example is just showing the usage of "PrivateLinkHubPrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapsePrivateLinkResource created on azure
// for more information of creating SynapsePrivateLinkResource, please refer to the document of SynapsePrivateLinkResource
string subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
string resourceGroupName = "ExampleResourceGroup";
string privateLinkHubName = "ExamplePrivateLinkHub";
string privateLinkResourceName = "sql";
ResourceIdentifier synapsePrivateLinkResourceId = SynapsePrivateLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateLinkHubName, privateLinkResourceName);
SynapsePrivateLinkResource synapsePrivateLinkResource = client.GetSynapsePrivateLinkResource(synapsePrivateLinkResourceId);

// invoke the operation
SynapsePrivateLinkResource result = await synapsePrivateLinkResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapsePrivateLinkResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
