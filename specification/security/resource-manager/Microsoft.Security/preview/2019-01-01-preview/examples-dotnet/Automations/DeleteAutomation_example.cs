using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/DeleteAutomation_example.json
// this example is just showing the usage of "Automations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityAutomationResource created on azure
// for more information of creating SecurityAutomationResource, please refer to the document of SecurityAutomationResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "myRg";
string automationName = "myAutomationName";
ResourceIdentifier securityAutomationResourceId = SecurityAutomationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationName);
SecurityAutomationResource securityAutomation = client.GetSecurityAutomationResource(securityAutomationResourceId);

// invoke the operation
await securityAutomation.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
