using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PrefixListLocalRulestack_Get_MaximumSet_Gen.json
// this example is just showing the usage of "PrefixListLocalRulestack_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LocalRulestackPrefixResource created on azure
// for more information of creating LocalRulestackPrefixResource, please refer to the document of LocalRulestackPrefixResource
string subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
string resourceGroupName = "rgopenapi";
string localRulestackName = "lrs1";
string name = "armid1";
ResourceIdentifier localRulestackPrefixResourceId = LocalRulestackPrefixResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, localRulestackName, name);
LocalRulestackPrefixResource localRulestackPrefix = client.GetLocalRulestackPrefixResource(localRulestackPrefixResourceId);

// invoke the operation
LocalRulestackPrefixResource result = await localRulestackPrefix.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LocalRulestackPrefixData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
