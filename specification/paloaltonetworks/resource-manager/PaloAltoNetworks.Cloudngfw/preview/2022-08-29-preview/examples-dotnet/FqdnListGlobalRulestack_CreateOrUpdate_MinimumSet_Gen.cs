using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/FqdnListGlobalRulestack_CreateOrUpdate_MinimumSet_Gen.json
// this example is just showing the usage of "FqdnListGlobalRulestack_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GlobalRulestackResource created on azure
// for more information of creating GlobalRulestackResource, please refer to the document of GlobalRulestackResource
string globalRulestackName = "praval";
ResourceIdentifier globalRulestackResourceId = GlobalRulestackResource.CreateResourceIdentifier(globalRulestackName);
GlobalRulestackResource globalRulestackResource = client.GetGlobalRulestackResource(globalRulestackResourceId);

// get the collection of this FqdnListGlobalRulestackResource
FqdnListGlobalRulestackResourceCollection collection = globalRulestackResource.GetFqdnListGlobalRulestackResources();

// invoke the operation
string name = "armid1";
FqdnListGlobalRulestackResourceData data = new FqdnListGlobalRulestackResourceData(new string[]
{
"string1","string2"
});
ArmOperation<FqdnListGlobalRulestackResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
FqdnListGlobalRulestackResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FqdnListGlobalRulestackResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
