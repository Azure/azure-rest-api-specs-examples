using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/MonitoringSettings_UpdatePut.json
// this example is just showing the usage of "MonitoringSettings_UpdatePut" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformMonitoringSettingResource created on azure
// for more information of creating AppPlatformMonitoringSettingResource, please refer to the document of AppPlatformMonitoringSettingResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
ResourceIdentifier appPlatformMonitoringSettingResourceId = AppPlatformMonitoringSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
AppPlatformMonitoringSettingResource appPlatformMonitoringSetting = client.GetAppPlatformMonitoringSettingResource(appPlatformMonitoringSettingResourceId);

// invoke the operation
AppPlatformMonitoringSettingData data = new AppPlatformMonitoringSettingData
{
    Properties = new AppPlatformMonitoringSettingProperties
    {
        IsTraceEnabled = true,
        AppInsightsInstrumentationKey = "00000000-0000-0000-0000-000000000000",
        AppInsightsSamplingRate = 10,
    },
};
ArmOperation<AppPlatformMonitoringSettingResource> lro = await appPlatformMonitoringSetting.CreateOrUpdateAsync(WaitUntil.Completed, data);
AppPlatformMonitoringSettingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppPlatformMonitoringSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
