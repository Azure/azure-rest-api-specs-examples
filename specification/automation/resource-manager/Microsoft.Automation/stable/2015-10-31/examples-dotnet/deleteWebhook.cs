using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/deleteWebhook.json
// this example is just showing the usage of "Webhook_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationWebhookResource created on azure
// for more information of creating AutomationWebhookResource, please refer to the document of AutomationWebhookResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount33";
string webhookName = "TestWebhook";
ResourceIdentifier automationWebhookResourceId = AutomationWebhookResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, webhookName);
AutomationWebhookResource automationWebhook = client.GetAutomationWebhookResource(automationWebhookResourceId);

// invoke the operation
await automationWebhook.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
