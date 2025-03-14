using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getConnectionType.json
// this example is just showing the usage of "ConnectionType_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationConnectionTypeResource created on azure
// for more information of creating AutomationConnectionTypeResource, please refer to the document of AutomationConnectionTypeResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount22";
string connectionTypeName = "myCT";
ResourceIdentifier automationConnectionTypeResourceId = AutomationConnectionTypeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, connectionTypeName);
AutomationConnectionTypeResource automationConnectionType = client.GetAutomationConnectionTypeResource(automationConnectionTypeResourceId);

// invoke the operation
AutomationConnectionTypeResource result = await automationConnectionType.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationConnectionTypeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
