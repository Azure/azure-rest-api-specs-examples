using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2025-04-15/examples/RuleSets_Get.json
// this example is just showing the usage of "FrontDoorRuleSets_Get" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
FrontDoorRuleSetResource result = await frontDoorRuleSet.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorRuleSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
