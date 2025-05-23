using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2022-02-22/examples/updateHybridRunbookWorkerGroup.json
// this example is just showing the usage of "HybridRunbookWorkerGroup_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridRunbookWorkerGroupResource created on azure
// for more information of creating HybridRunbookWorkerGroupResource, please refer to the document of HybridRunbookWorkerGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "testaccount";
string hybridRunbookWorkerGroupName = "TestHybridGroup";
ResourceIdentifier hybridRunbookWorkerGroupResourceId = HybridRunbookWorkerGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName);
HybridRunbookWorkerGroupResource hybridRunbookWorkerGroup = client.GetHybridRunbookWorkerGroupResource(hybridRunbookWorkerGroupResourceId);

// invoke the operation
HybridRunbookWorkerGroupCreateOrUpdateContent content = new HybridRunbookWorkerGroupCreateOrUpdateContent
{
    CredentialName = "myRunAsCredentialUpdatedName",
};
HybridRunbookWorkerGroupResource result = await hybridRunbookWorkerGroup.UpdateAsync(content);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridRunbookWorkerGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
