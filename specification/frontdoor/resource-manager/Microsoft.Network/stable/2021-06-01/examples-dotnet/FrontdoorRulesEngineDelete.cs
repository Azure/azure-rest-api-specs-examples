using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorRulesEngineDelete.json
// this example is just showing the usage of "RulesEngines_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorRulesEngineResource created on azure
// for more information of creating FrontDoorRulesEngineResource, please refer to the document of FrontDoorRulesEngineResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string frontDoorName = "frontDoor1";
string rulesEngineName = "rulesEngine1";
ResourceIdentifier frontDoorRulesEngineResourceId = FrontDoorRulesEngineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, frontDoorName, rulesEngineName);
FrontDoorRulesEngineResource frontDoorRulesEngine = client.GetFrontDoorRulesEngineResource(frontDoorRulesEngineResourceId);

// invoke the operation
await frontDoorRulesEngine.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
