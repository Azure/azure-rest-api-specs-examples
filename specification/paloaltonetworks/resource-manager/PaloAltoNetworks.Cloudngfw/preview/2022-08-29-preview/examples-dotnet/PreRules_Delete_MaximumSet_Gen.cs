using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/PreRules_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "PreRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PreRulesResource created on azure
// for more information of creating PreRulesResource, please refer to the document of PreRulesResource
string globalRulestackName = "lrs1";
string priority = "1";
ResourceIdentifier preRulesResourceId = PreRulesResource.CreateResourceIdentifier(globalRulestackName, priority);
PreRulesResource preRulesResource = client.GetPreRulesResource(preRulesResourceId);

// invoke the operation
await preRulesResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
