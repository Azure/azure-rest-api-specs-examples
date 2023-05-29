using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/getDscConfiguration.json
// this example is just showing the usage of "DscConfiguration_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DscConfigurationResource created on azure
// for more information of creating DscConfigurationResource, please refer to the document of DscConfigurationResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount33";
string configurationName = "TemplateBasic";
ResourceIdentifier dscConfigurationResourceId = DscConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, configurationName);
DscConfigurationResource dscConfiguration = client.GetDscConfigurationResource(dscConfigurationResourceId);

// invoke the operation
DscConfigurationResource result = await dscConfiguration.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DscConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
