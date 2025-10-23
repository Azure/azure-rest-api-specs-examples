using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PolicyInsights;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/Remediations_ListResourceScope_WithQuery.json
// this example is just showing the usage of "Remediations_ListForResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this PolicyRemediationResource
string resourceId = "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storAc1";
PolicyRemediationCollection collection = client.GetPolicyRemediations(new ResourceIdentifier(resourceId));

// invoke the operation and iterate over the result
await foreach (PolicyRemediationResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PolicyRemediationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
