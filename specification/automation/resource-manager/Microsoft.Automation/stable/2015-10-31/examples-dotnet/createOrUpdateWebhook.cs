using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/createOrUpdateWebhook.json
// this example is just showing the usage of "Webhook_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationAccountResource created on azure
// for more information of creating AutomationAccountResource, please refer to the document of AutomationAccountResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount33";
ResourceIdentifier automationAccountResourceId = AutomationAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName);
AutomationAccountResource automationAccount = client.GetAutomationAccountResource(automationAccountResourceId);

// get the collection of this AutomationWebhookResource
AutomationWebhookCollection collection = automationAccount.GetAutomationWebhooks();

// invoke the operation
string webhookName = "TestWebhook";
AutomationWebhookCreateOrUpdateContent content = new AutomationWebhookCreateOrUpdateContent("TestWebhook")
{
    IsEnabled = true,
    Uri = new Uri("<uri>"),
    ExpireOn = DateTimeOffset.Parse("2018-03-29T22:18:13.7002872Z"),
    RunbookName = "TestRunbook",
};
ArmOperation<AutomationWebhookResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, webhookName, content);
AutomationWebhookResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationWebhookData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
