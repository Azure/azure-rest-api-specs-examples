using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/FqdnListGlobalRulestack_Delete_MinimumSet_Gen.json
// this example is just showing the usage of "FqdnListGlobalRulestack_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GlobalRulestackFqdnResource created on azure
// for more information of creating GlobalRulestackFqdnResource, please refer to the document of GlobalRulestackFqdnResource
string globalRulestackName = "praval";
string name = "armid1";
ResourceIdentifier globalRulestackFqdnResourceId = GlobalRulestackFqdnResource.CreateResourceIdentifier(globalRulestackName, name);
GlobalRulestackFqdnResource globalRulestackFqdn = client.GetGlobalRulestackFqdnResource(globalRulestackFqdnResourceId);

// invoke the operation
await globalRulestackFqdn.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
