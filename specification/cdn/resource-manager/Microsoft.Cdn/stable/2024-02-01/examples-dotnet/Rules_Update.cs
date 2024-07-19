using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Rules_Update.json
// this example is just showing the usage of "FrontDoorRules_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorRuleResource created on azure
// for more information of creating FrontDoorRuleResource, please refer to the document of FrontDoorRuleResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string ruleSetName = "ruleSet1";
string ruleName = "rule1";
ResourceIdentifier frontDoorRuleResourceId = FrontDoorRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, ruleSetName, ruleName);
FrontDoorRuleResource frontDoorRule = client.GetFrontDoorRuleResource(frontDoorRuleResourceId);

// invoke the operation
FrontDoorRulePatch patch = new FrontDoorRulePatch()
{
    Order = 1,
    Actions =
    {
    new DeliveryRuleResponseHeaderAction(new HeaderActionProperties(HeaderActionType.HeaderAction,HeaderAction.Overwrite,"X-CDN")
    {
    Value = "MSFT",
    })
    },
};
ArmOperation<FrontDoorRuleResource> lro = await frontDoorRule.UpdateAsync(WaitUntil.Completed, patch);
FrontDoorRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
