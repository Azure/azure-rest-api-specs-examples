using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/GetAutomationResourceGroup_example.json
// this example is just showing the usage of "Automations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityAutomationResource created on azure
// for more information of creating SecurityAutomationResource, please refer to the document of SecurityAutomationResource
string subscriptionId = "a5caac9c-5c04-49af-b3d0-e204f40345d5";
string resourceGroupName = "exampleResourceGroup";
string automationName = "exampleAutomation";
ResourceIdentifier securityAutomationResourceId = SecurityAutomationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationName);
SecurityAutomationResource securityAutomation = client.GetSecurityAutomationResource(securityAutomationResourceId);

// invoke the operation
SecurityAutomationResource result = await securityAutomation.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityAutomationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
