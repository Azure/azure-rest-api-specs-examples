using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateDscNodeConfiguration.json
// this example is just showing the usage of "DscNodeConfiguration_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DscNodeConfigurationResource created on azure
// for more information of creating DscNodeConfigurationResource, please refer to the document of DscNodeConfigurationResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount20";
string nodeConfigurationName = "configName.nodeConfigName";
ResourceIdentifier dscNodeConfigurationResourceId = DscNodeConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, nodeConfigurationName);
DscNodeConfigurationResource dscNodeConfiguration = client.GetDscNodeConfigurationResource(dscNodeConfigurationResourceId);

// invoke the operation
DscNodeConfigurationCreateOrUpdateContent content = new DscNodeConfigurationCreateOrUpdateContent()
{
    Name = "configName.nodeConfigName",
    Source = new AutomationContentSource()
    {
        Hash = new AutomationContentHash("sha256", "6DE256A57F01BFA29B88696D5E77A383D6E61484C7686E8DB955FA10ACE9FFE5"),
        SourceType = AutomationContentSourceType.EmbeddedContent,
        Value = "\r\ninstance of MSFT_RoleResource as $MSFT_RoleResource1ref\r\n{\r\nResourceID = \"[WindowsFeature]IIS\";\r\n Ensure = \"Present\";\r\n SourceInfo = \"::3::32::WindowsFeature\";\r\n Name = \"Web-Server\";\r\n ModuleName = \"PsDesiredStateConfiguration\";\r\n\r\nModuleVersion = \"1.0\";\r\r\n ConfigurationName = \"configName\";\r\r\n};\r\ninstance of OMI_ConfigurationDocument\r\n\r\r\n                    {\r\n Version=\"2.0.0\";\r\n \r\r\n                        MinimumCompatibleVersion = \"1.0.0\";\r\n \r\r\n                        CompatibleVersionAdditionalProperties= {\"Omi_BaseResource:ConfigurationName\"};\r\n \r\r\n                        Author=\"weijiel\";\r\n \r\r\n                        GenerationDate=\"03/30/2017 13:40:25\";\r\n \r\r\n                        GenerationHost=\"TEST-BACKEND\";\r\n \r\r\n                        Name=\"configName\";\r\n\r\r\n                    };\r\n",
        Version = "1.0",
    },
    ConfigurationName = "configName",
    IsIncrementNodeConfigurationBuildRequired = true,
};
await dscNodeConfiguration.UpdateAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
