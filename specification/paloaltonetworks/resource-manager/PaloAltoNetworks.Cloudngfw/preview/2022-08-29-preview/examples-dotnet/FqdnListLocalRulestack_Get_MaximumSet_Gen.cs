using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/FqdnListLocalRulestack_Get_MaximumSet_Gen.json
// this example is just showing the usage of "FqdnListLocalRulestack_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FqdnListLocalRulestackResource created on azure
// for more information of creating FqdnListLocalRulestackResource, please refer to the document of FqdnListLocalRulestackResource
string subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
string resourceGroupName = "rgopenapi";
string localRulestackName = "lrs1";
string name = "armid1";
ResourceIdentifier fqdnListLocalRulestackResourceId = FqdnListLocalRulestackResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, localRulestackName, name);
FqdnListLocalRulestackResource fqdnListLocalRulestackResource = client.GetFqdnListLocalRulestackResource(fqdnListLocalRulestackResourceId);

// invoke the operation
FqdnListLocalRulestackResource result = await fqdnListLocalRulestackResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FqdnListLocalRulestackResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
