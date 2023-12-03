using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkFunction;
using Azure.ResourceManager.NetworkFunction.Models;

// Generated from example definition: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/CollectorPolicyGet.json
// this example is just showing the usage of "CollectorPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AzureTrafficCollectorResource created on azure
// for more information of creating AzureTrafficCollectorResource, please refer to the document of AzureTrafficCollectorResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string azureTrafficCollectorName = "atc";
ResourceIdentifier azureTrafficCollectorResourceId = AzureTrafficCollectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, azureTrafficCollectorName);
AzureTrafficCollectorResource azureTrafficCollector = client.GetAzureTrafficCollectorResource(azureTrafficCollectorResourceId);

// get the collection of this CollectorPolicyResource
CollectorPolicyCollection collection = azureTrafficCollector.GetCollectorPolicies();

// invoke the operation
string collectorPolicyName = "cp1";
NullableResponse<CollectorPolicyResource> response = await collection.GetIfExistsAsync(collectorPolicyName);
CollectorPolicyResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CollectorPolicyData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
