using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdatePython2Package.json
// this example is just showing the usage of "Python2Package_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this AutomationAccountPython2PackageResource
AutomationAccountPython2PackageCollection collection = automationAccount.GetAutomationAccountPython2Packages();

// invoke the operation
string packageName = "OmsCompositeResources";
AutomationAccountPython2PackageCreateOrUpdateContent content = new AutomationAccountPython2PackageCreateOrUpdateContent(new AutomationContentLink()
{
    Uri = new Uri("https://teststorage.blob.core.windows.net/dsccomposite/OmsCompositeResources.zip"),
    ContentHash = new AutomationContentHash("sha265", "07E108A962B81DD9C9BAA89BB47C0F6EE52B29E83758B07795E408D258B2B87A"),
    Version = "1.0.0.0",
})
{
    Tags =
    {
    },
};
ArmOperation<AutomationAccountPython2PackageResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, packageName, content);
AutomationAccountPython2PackageResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationModuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
