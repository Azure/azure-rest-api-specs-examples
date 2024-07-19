using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Rules_Create.json
// this example is just showing the usage of "FrontDoorRules_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorRuleSetResource created on azure
// for more information of creating FrontDoorRuleSetResource, please refer to the document of FrontDoorRuleSetResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string ruleSetName = "ruleSet1";
ResourceIdentifier frontDoorRuleSetResourceId = FrontDoorRuleSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, ruleSetName);
FrontDoorRuleSetResource frontDoorRuleSet = client.GetFrontDoorRuleSetResource(frontDoorRuleSetResourceId);

// get the collection of this FrontDoorRuleResource
FrontDoorRuleCollection collection = frontDoorRuleSet.GetFrontDoorRules();

// invoke the operation
string ruleName = "rule1";
FrontDoorRuleData data = new FrontDoorRuleData()
{
    Order = 1,
    Conditions =
    {
    new DeliveryRuleRequestMethodCondition(new RequestMethodMatchCondition(RequestMethodMatchConditionType.RequestMethodCondition,RequestMethodOperator.Equal)
    {
    NegateCondition = false,
    MatchValues =
    {
    RequestMethodMatchConditionMatchValue.Get
    },
    })
    },
    Actions =
    {
    new DeliveryRuleResponseHeaderAction(new HeaderActionProperties(HeaderActionType.HeaderAction,HeaderAction.Overwrite,"X-CDN")
    {
    Value = "MSFT",
    })
    },
};
ArmOperation<FrontDoorRuleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, ruleName, data);
FrontDoorRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
