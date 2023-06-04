using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Quota;

// Generated from example definition: specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getComputeUsages.json
// this example is just showing the usage of "Usages_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this CurrentUsagesBaseResource
string scope = "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
CurrentUsagesBaseCollection collection = client.GetCurrentUsagesBases(scopeId);

// invoke the operation and iterate over the result
await foreach (CurrentUsagesBaseResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CurrentUsagesBaseData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
