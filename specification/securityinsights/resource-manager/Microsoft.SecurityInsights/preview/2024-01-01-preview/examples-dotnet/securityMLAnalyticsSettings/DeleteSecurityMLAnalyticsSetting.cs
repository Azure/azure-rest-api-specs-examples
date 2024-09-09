using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/securityMLAnalyticsSettings/DeleteSecurityMLAnalyticsSetting.json
// this example is just showing the usage of "SecurityMLAnalyticsSettings_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityMLAnalyticsSettingResource created on azure
// for more information of creating SecurityMLAnalyticsSettingResource, please refer to the document of SecurityMLAnalyticsSettingResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string settingsResourceName = "f209187f-1d17-4431-94af-c141bf5f23db";
ResourceIdentifier securityMLAnalyticsSettingResourceId = SecurityMLAnalyticsSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, settingsResourceName);
SecurityMLAnalyticsSettingResource securityMLAnalyticsSetting = client.GetSecurityMLAnalyticsSettingResource(securityMLAnalyticsSettingResourceId);

// invoke the operation
await securityMLAnalyticsSetting.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
