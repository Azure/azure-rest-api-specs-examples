using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/PrefixListGlobalRulestack_CreateOrUpdate_MinimumSet_Gen.json
// this example is just showing the usage of "PrefixListGlobalRulestack_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GlobalRulestackPrefixResource created on azure
// for more information of creating GlobalRulestackPrefixResource, please refer to the document of GlobalRulestackPrefixResource
string globalRulestackName = "praval";
string name = "armid1";
ResourceIdentifier globalRulestackPrefixResourceId = GlobalRulestackPrefixResource.CreateResourceIdentifier(globalRulestackName, name);
GlobalRulestackPrefixResource globalRulestackPrefix = client.GetGlobalRulestackPrefixResource(globalRulestackPrefixResourceId);

// invoke the operation
GlobalRulestackPrefixData data = new GlobalRulestackPrefixData(new string[]
{
"1.0.0.0/24"
});
ArmOperation<GlobalRulestackPrefixResource> lro = await globalRulestackPrefix.UpdateAsync(WaitUntil.Completed, data);
GlobalRulestackPrefixResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GlobalRulestackPrefixData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
