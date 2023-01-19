using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Cdn;
using Azure.ResourceManager.Cdn.Models;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Get.json
// this example is just showing the usage of "FrontDoorRules_Get" operation, for the dependent resources, they will have to be created separately.

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
bool result = await collection.ExistsAsync(ruleName);

Console.WriteLine($"Succeeded: {result}");
