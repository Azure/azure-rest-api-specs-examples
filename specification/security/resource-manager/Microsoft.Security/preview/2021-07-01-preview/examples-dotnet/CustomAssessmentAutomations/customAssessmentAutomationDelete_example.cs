using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomAssessmentAutomations/customAssessmentAutomationDelete_example.json
// this example is just showing the usage of "CustomAssessmentAutomations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CustomAssessmentAutomationResource created on azure
// for more information of creating CustomAssessmentAutomationResource, please refer to the document of CustomAssessmentAutomationResource
string subscriptionId = "e5d1b86c-3051-44d5-8802-aa65d45a279b";
string resourceGroupName = "TestResourceGroup";
string customAssessmentAutomationName = "MyCustomAssessmentAutomation";
ResourceIdentifier customAssessmentAutomationResourceId = CustomAssessmentAutomationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, customAssessmentAutomationName);
CustomAssessmentAutomationResource customAssessmentAutomation = client.GetCustomAssessmentAutomationResource(customAssessmentAutomationResourceId);

// invoke the operation
await customAssessmentAutomation.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
