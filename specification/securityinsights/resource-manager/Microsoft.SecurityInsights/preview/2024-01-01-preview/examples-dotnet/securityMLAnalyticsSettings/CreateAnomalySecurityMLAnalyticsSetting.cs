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

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/securityMLAnalyticsSettings/CreateAnomalySecurityMLAnalyticsSetting.json
// this example is just showing the usage of "SecurityMLAnalyticsSettings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
SecurityMLAnalyticsSettingData data = new AnomalySecurityMLAnalyticsSettings()
{
    Description = "When account logs from a source region that has rarely been logged in from during the last 14 days, an anomaly is triggered.",
    DisplayName = "Login from unusual region",
    IsEnabled = true,
    RequiredDataConnectors =
    {
    new SecurityMLAnalyticsSettingsDataSource()
    {
    ConnectorId = "AWS",
    DataTypes =
    {
    "AWSCloudTrail"
    },
    }
    },
    Tactics =
    {
    SecurityInsightsAttackTactic.Exfiltration,SecurityInsightsAttackTactic.CommandAndControl
    },
    Techniques =
    {
    "T1037","T1021"
    },
    AnomalyVersion = "1.0.5",
    CustomizableObservations = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
        ["multiSelectObservations"] = null,
        ["prioritizeExcludeObservations"] = null,
        ["singleSelectObservations"] = new object[] { new Dictionary<string, object>()
        {
        ["name"] = "Device vendor",
        ["description"] = "Select device vendor of network connection logs from CommonSecurityLog",
        ["rerun"] = "RerunAlways",
        ["sequenceNumber"] = "1",
        ["supportedValues"] = new object[] { "Palo Alto Networks", "Fortinet", "Check Point" },
        ["supportedValuesKql"] = null,
        ["value"] = new object[] { "Palo Alto Networks" },
        ["valuesKql"] = null} },
        ["singleValueObservations"] = null,
        ["thresholdObservations"] = new object[] { new Dictionary<string, object>()
        {
        ["name"] = "Daily data transfer threshold in MB",
        ["description"] = "Suppress anomalies when daily data transfered (in MB) per hour is less than the chosen value",
        ["maximum"] = "100",
        ["minimum"] = "1",
        ["rerun"] = "RerunAlways",
        ["sequenceNumber"] = "1",
        ["value"] = "25"}, new Dictionary<string, object>()
        {
        ["name"] = "Number of standard deviations",
        ["description"] = "Triggers anomalies when number of standard deviations is greater than the chosen value",
        ["maximum"] = "10",
        ["minimum"] = "2",
        ["rerun"] = "RerunAlways",
        ["sequenceNumber"] = "2",
        ["value"] = "3"} }
    }),
    Frequency = XmlConvert.ToTimeSpan("PT1H"),
    SettingsStatus = AnomalySecurityMLAnalyticsSettingsStatus.Production,
    IsDefaultSettings = true,
    AnomalySettingsVersion = 0,
    SettingsDefinitionId = Guid.Parse("f209187f-1d17-4431-94af-c141bf5f23db"),
    ETag = new ETag("\"260090e2-0000-0d00-0000-5d6fb8670000\""),
};
ArmOperation<SecurityMLAnalyticsSettingResource> lro = await securityMLAnalyticsSetting.UpdateAsync(WaitUntil.Completed, data);
SecurityMLAnalyticsSettingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityMLAnalyticsSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
