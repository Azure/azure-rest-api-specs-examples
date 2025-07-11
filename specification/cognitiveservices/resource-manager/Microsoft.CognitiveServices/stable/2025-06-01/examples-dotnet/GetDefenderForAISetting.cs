using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/GetDefenderForAISetting.json
// this example is just showing the usage of "DefenderForAISettings_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DefenderForAISettingResource created on azure
// for more information of creating DefenderForAISettingResource, please refer to the document of DefenderForAISettingResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resourceGroupName";
string accountName = "accountName";
string defenderForAISettingName = "Default";
ResourceIdentifier defenderForAISettingResourceId = DefenderForAISettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, defenderForAISettingName);
DefenderForAISettingResource defenderForAISetting = client.GetDefenderForAISettingResource(defenderForAISettingResourceId);

// invoke the operation
DefenderForAISettingResource result = await defenderForAISetting.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DefenderForAISettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
