using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/GlobalRulestack_listAppIds_MaximumSet_Gen.json
// this example is just showing the usage of "GlobalRulestack_listAppIds" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GlobalRulestackResource created on azure
// for more information of creating GlobalRulestackResource, please refer to the document of GlobalRulestackResource
string globalRulestackName = "praval";
ResourceIdentifier globalRulestackResourceId = GlobalRulestackResource.CreateResourceIdentifier(globalRulestackName);
GlobalRulestackResource globalRulestack = client.GetGlobalRulestackResource(globalRulestackResourceId);

// invoke the operation and iterate over the result
string appIdVersion = "8543";
string appPrefix = "pref";
string skip = "a6a321";
int? top = 20;
await foreach (string item in globalRulestack.GetAppIdsAsync(appIdVersion: appIdVersion, appPrefix: appPrefix, skip: skip, top: top))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
