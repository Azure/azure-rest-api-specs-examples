using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateConnection.json
// this example is just showing the usage of "Connection_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationAccountResource created on azure
// for more information of creating AutomationAccountResource, please refer to the document of AutomationAccountResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount28";
ResourceIdentifier automationAccountResourceId = AutomationAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName);
AutomationAccountResource automationAccount = client.GetAutomationAccountResource(automationAccountResourceId);

// get the collection of this AutomationConnectionResource
AutomationConnectionCollection collection = automationAccount.GetAutomationConnections();

// invoke the operation
string connectionName = "mysConnection";
AutomationConnectionCreateOrUpdateContent content = new AutomationConnectionCreateOrUpdateContent("mysConnection", new ConnectionTypeAssociationProperty()
{
    Name = "Azure",
})
{
    Description = "my description goes here",
    FieldDefinitionValues =
    {
    ["AutomationCertificateName"] = "mysCertificateName",
    ["SubscriptionID"] = "subid",
    },
};
ArmOperation<AutomationConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, connectionName, content);
AutomationConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
