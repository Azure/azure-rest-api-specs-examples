using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/createOrUpdateAutomationAccount.json
// this example is just showing the usage of "AutomationAccount_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AutomationAccountResource
AutomationAccountCollection collection = resourceGroupResource.GetAutomationAccounts();

// invoke the operation
string automationAccountName = "myAutomationAccount9";
AutomationAccountCreateOrUpdateContent content = new AutomationAccountCreateOrUpdateContent
{
    Name = "myAutomationAccount9",
    Location = new AzureLocation("East US 2"),
    Sku = new AutomationSku(AutomationSkuName.Free),
};
ArmOperation<AutomationAccountResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, automationAccountName, content);
AutomationAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
