using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getCredential.json
// this example is just showing the usage of "Credential_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationCredentialResource created on azure
// for more information of creating AutomationCredentialResource, please refer to the document of AutomationCredentialResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount18";
string credentialName = "myCredential";
ResourceIdentifier automationCredentialResourceId = AutomationCredentialResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, credentialName);
AutomationCredentialResource automationCredential = client.GetAutomationCredentialResource(automationCredentialResourceId);

// invoke the operation
AutomationCredentialResource result = await automationCredential.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationCredentialData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
