using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/LocalRulestacks_listCountries_MaximumSet_Gen.json
// this example is just showing the usage of "LocalRulestacks_ListCountries" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LocalRulestackResource created on azure
// for more information of creating LocalRulestackResource, please refer to the document of LocalRulestackResource
string subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
string resourceGroupName = "rgopenapi";
string localRulestackName = "lrs1";
ResourceIdentifier localRulestackResourceId = LocalRulestackResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, localRulestackName);
LocalRulestackResource localRulestack = client.GetLocalRulestackResource(localRulestackResourceId);

// invoke the operation and iterate over the result
string skip = "a6a321";
int? top = 20;
await foreach (RulestackCountry item in localRulestack.GetCountriesAsync(skip: skip, top: top))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
