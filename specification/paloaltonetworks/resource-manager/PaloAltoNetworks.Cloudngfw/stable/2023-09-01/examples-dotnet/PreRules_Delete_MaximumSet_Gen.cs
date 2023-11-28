using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "PreRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PreRulestackRuleResource created on azure
// for more information of creating PreRulestackRuleResource, please refer to the document of PreRulestackRuleResource
string globalRulestackName = "lrs1";
string priority = "1";
ResourceIdentifier preRulestackRuleResourceId = PreRulestackRuleResource.CreateResourceIdentifier(globalRulestackName, priority);
PreRulestackRuleResource preRulestackRule = client.GetPreRulestackRuleResource(preRulestackRuleResourceId);

// invoke the operation
await preRulestackRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
