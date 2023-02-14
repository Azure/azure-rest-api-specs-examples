using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/ValidateAutomation_example.json
// this example is just showing the usage of "Automations_Validate" operation, for the dependent resources, they will have to be created separately.

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
    Description = "An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment of type customAssessment",
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
    RuleSets =
    {
    new SecurityAutomationRuleSet()
    {
    Rules =
    {
    new SecurityAutomationTriggeringRule()
    {
    PropertyJPath = "$.Entity.AssessmentType",
    PropertyType = AutomationTriggeringRulePropertyType.String,
    ExpectedValue = "customAssessment",
    Operator = AutomationTriggeringRuleOperator.EqualsValue,
    }
    },
    }
    },
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
    Tags =
    {
    },
};
SecurityAutomationValidationStatus result = await securityAutomation.ValidateAsync(data);

Console.WriteLine($"Succeeded: {result}");
