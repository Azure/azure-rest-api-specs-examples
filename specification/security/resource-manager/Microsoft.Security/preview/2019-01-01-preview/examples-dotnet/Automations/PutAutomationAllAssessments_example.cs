using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/PutAutomationAllAssessments_example.json
// this example is just showing the usage of "Automations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
SecurityAutomationData data = new SecurityAutomationData(new AzureLocation("Central US"))
{
    Description = "An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment",
    IsEnabled = true,
    Scopes =
    {
    new SecurityAutomationScope()
    {
    Description = "A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5",
    ScopePath = "/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup",
    }
    },
    Sources =
    {
    new SecurityAutomationSource()
    {
    EventSource = SecurityEventSource.Assessments,
    }
    },
    Actions =
    {
    new SecurityAutomationActionLogicApp()
    {
    LogicAppResourceId = new ResourceIdentifier("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
    Uri = new Uri("https://exampleTriggerUri1.com"),
    }
    },
    ETag = new ETag("etag value (must be supplied for update)"),
    Tags =
    {
    },
};
ArmOperation<SecurityAutomationResource> lro = await securityAutomation.UpdateAsync(WaitUntil.Completed, data);
SecurityAutomationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityAutomationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
