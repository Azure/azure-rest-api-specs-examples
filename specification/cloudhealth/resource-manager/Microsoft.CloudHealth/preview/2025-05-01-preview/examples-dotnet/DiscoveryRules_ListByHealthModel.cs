using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CloudHealth.Models;
using Azure.ResourceManager.CloudHealth;

// Generated from example definition: 2025-05-01-preview/DiscoveryRules_ListByHealthModel.json
// this example is just showing the usage of "DiscoveryRule_ListByHealthModel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthModelResource created on azure
// for more information of creating HealthModelResource, please refer to the document of HealthModelResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "my-resource-group";
string healthModelName = "my-health-model";
ResourceIdentifier healthModelResourceId = HealthModelResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, healthModelName);
HealthModelResource healthModel = client.GetHealthModelResource(healthModelResourceId);

// get the collection of this HealthModelDiscoveryRuleResource
HealthModelDiscoveryRuleCollection collection = healthModel.GetHealthModelDiscoveryRules();

// invoke the operation and iterate over the result
await foreach (HealthModelDiscoveryRuleResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HealthModelDiscoveryRuleData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
