using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay.Models;
using Azure.ResourceManager.Relay;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayGet.json
// this example is just showing the usage of "WCFRelays_Get" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this WcfRelayResource created on azure
// for more information of creating WcfRelayResource, please refer to the document of WcfRelayResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "resourcegroup";
string namespaceName = "example-RelayNamespace-9953";
string relayName = "example-Relay-Wcf-1194";
ResourceIdentifier wcfRelayResourceId = WcfRelayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, relayName);
WcfRelayResource wcfRelay = client.GetWcfRelayResource(wcfRelayResourceId);

// invoke the operation
WcfRelayResource result = await wcfRelay.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WcfRelayData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
