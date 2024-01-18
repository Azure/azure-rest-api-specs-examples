using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/GetAdaptiveNetworkHardening_example.json
// this example is just showing the usage of "AdaptiveNetworkHardenings_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AdaptiveNetworkHardeningResource created on azure
// for more information of creating AdaptiveNetworkHardeningResource, please refer to the document of AdaptiveNetworkHardeningResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "rg1";
string resourceNamespace = "Microsoft.Compute";
string resourceType = "virtualMachines";
string resourceName = "vm1";
string adaptiveNetworkHardeningResourceName = "default";
ResourceIdentifier adaptiveNetworkHardeningResourceId = AdaptiveNetworkHardeningResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceNamespace, resourceType, resourceName, adaptiveNetworkHardeningResourceName);
AdaptiveNetworkHardeningResource adaptiveNetworkHardening = client.GetAdaptiveNetworkHardeningResource(adaptiveNetworkHardeningResourceId);

// invoke the operation
AdaptiveNetworkHardeningResource result = await adaptiveNetworkHardening.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AdaptiveNetworkHardeningData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
