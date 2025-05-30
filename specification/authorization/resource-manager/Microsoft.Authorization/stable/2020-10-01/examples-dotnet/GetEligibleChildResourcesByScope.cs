using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Authorization.Models;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetEligibleChildResourcesByScope.json
// this example is just showing the usage of "EligibleChildResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// invoke the operation and iterate over the result
ResourceIdentifier scope = new ResourceIdentifier("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f");
string filter = "resourceType+eq+'resourcegroup'";
await foreach (EligibleChildResource item in client.GetEligibleChildResourcesAsync(scope, filter: filter))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
