using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkFunction;
using Azure.ResourceManager.NetworkFunction.Models;

// Generated from example definition: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/CollectorPolicyDelete.json
// this example is just showing the usage of "CollectorPolicies_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CollectorPolicyResource created on azure
// for more information of creating CollectorPolicyResource, please refer to the document of CollectorPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string azureTrafficCollectorName = "atc";
string collectorPolicyName = "cp1";
ResourceIdentifier collectorPolicyResourceId = CollectorPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, azureTrafficCollectorName, collectorPolicyName);
CollectorPolicyResource collectorPolicy = client.GetCollectorPolicyResource(collectorPolicyResourceId);

// invoke the operation
await collectorPolicy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
