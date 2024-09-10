using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/settings/DeleteEyesOnSetting.json
// this example is just showing the usage of "ProductSettings_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsSettingResource created on azure
// for more information of creating SecurityInsightsSettingResource, please refer to the document of SecurityInsightsSettingResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string settingsName = "EyesOn";
ResourceIdentifier securityInsightsSettingResourceId = SecurityInsightsSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, settingsName);
SecurityInsightsSettingResource securityInsightsSetting = client.GetSecurityInsightsSettingResource(securityInsightsSettingResourceId);

// invoke the operation
await securityInsightsSetting.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
