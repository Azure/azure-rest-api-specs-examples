using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2024-01-01/examples/NameSpaces/VirtualNetworkRule/SBNetworkRuleSetGet.json
// this example is just showing the usage of "Namespaces_GetNetworkRuleSet" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusNetworkRuleSetResource created on azure
// for more information of creating ServiceBusNetworkRuleSetResource, please refer to the document of ServiceBusNetworkRuleSetResource
string subscriptionId = "Subscription";
string resourceGroupName = "ResourceGroup";
string namespaceName = "sdk-Namespace-6019";
ResourceIdentifier serviceBusNetworkRuleSetResourceId = ServiceBusNetworkRuleSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
ServiceBusNetworkRuleSetResource serviceBusNetworkRuleSet = client.GetServiceBusNetworkRuleSetResource(serviceBusNetworkRuleSetResourceId);

// invoke the operation
ServiceBusNetworkRuleSetResource result = await serviceBusNetworkRuleSet.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceBusNetworkRuleSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
