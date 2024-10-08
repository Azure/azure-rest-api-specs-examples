using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/contentTemplates/InstallTemplate.json
// this example is just showing the usage of "ContentTemplate_Install" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsTemplateResource created on azure
// for more information of creating SecurityInsightsTemplateResource, please refer to the document of SecurityInsightsTemplateResource
string subscriptionId = "d0cfeab2-9ae0-4464-9919-dccaee2e48f0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string templateId = "str.azure-sentinel-solution-str";
ResourceIdentifier securityInsightsTemplateResourceId = SecurityInsightsTemplateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, templateId);
SecurityInsightsTemplateResource securityInsightsTemplate = client.GetSecurityInsightsTemplateResource(securityInsightsTemplateResourceId);

// invoke the operation
SecurityInsightsTemplateData data = new SecurityInsightsTemplateData()
{
    ContentId = "8365ebfe-a381-45b7-ad08-7d818070e11f",
    ContentProductId = "str.azure-sentinel-solution-str-ar-cbfe4fndz66bi",
    PackageVersion = "1.0.0",
    Version = "1.0.1",
    DisplayName = "API Protection workbook template",
    ContentKind = SecurityInsightsKind.AnalyticsRule,
    Source = new SecurityInsightsMetadataSource(SecurityInsightsSourceKind.Solution)
    {
        Name = "str",
        SourceId = "str.azure-sentinel-solution-str",
    },
    Author = new SecurityInsightsMetadataAuthor()
    {
        Name = "Microsoft",
        Email = "support@microsoft.com",
    },
    Support = new SecurityInsightsMetadataSupport(SecurityInsightsSupportTier.Microsoft)
    {
        Name = "Microsoft Corporation",
        Email = "support@microsoft.com",
        Link = "https://support.microsoft.com/",
    },
    PackageId = "str.azure-sentinel-solution-str",
    PackageKind = SecurityInsightsMetadataPackageKind.Solution,
    PackageName = "str",
    MainTemplate = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
        ["$schema"] = "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
        ["contentVersion"] = "1.0.1",
        ["resources"] = new object[] { new Dictionary<string, object>()
        {
        ["name"] = "8365ebfe-a381-45b7-ad08-7d818070e11f",
        ["type"] = "Microsoft.SecurityInsights/AlertRuleTemplates",
        ["apiVersion"] = "2022-04-01-preview",
        ["kind"] = "Scheduled",
        ["location"] = "[parameters('workspace-location')]",
        ["properties"] = new Dictionary<string, object>()
        {
        ["description"] = "Creates an incident when a large number of Critical/High severity CrowdStrike Falcon sensor detections is triggered by a single user",
        ["displayName"] = "Critical or High Severity Detections by User",
        ["enabled"] = "false",
        ["query"] = "...",
        ["queryFrequency"] = "PT1H",
        ["queryPeriod"] = "PT1H",
        ["severity"] = "High",
        ["status"] = "Available",
        ["suppressionDuration"] = "PT1H",
        ["suppressionEnabled"] = "false",
        ["triggerOperator"] = "GreaterThan",
        ["triggerThreshold"] = "0"}}, new Dictionary<string, object>()
        {
        ["name"] = "[concat(parameters('workspace'),'/Microsoft.SecurityInsights/',concat('AnalyticsRule-', last(split([resourceId('Microsoft.SecurityInsights/AlertRuleTemplates', 8365ebfe-a381-45b7-ad08-7d818070e11f)],'/'))))]",
        ["type"] = "Microsoft.OperationalInsights/workspaces/providers/metadata",
        ["apiVersion"] = "2022-01-01-preview",
        ["properties"] = new Dictionary<string, object>()
        {
        ["description"] = "CrowdStrike Falcon Endpoint Protection Analytics Rule 1",
        ["author"] = new Dictionary<string, object>()
        {
        ["name"] = "Microsoft",
        ["email"] = "support@microsoft.com"},
        ["contentId"] = "4465ebde-b381-45f7-ad08-7d818070a11c",
        ["kind"] = "AnalyticsRule",
        ["parentId"] = "[resourceId('Microsoft.SecurityInsights/AlertRuleTemplates', 8365ebfe-a381-45b7-ad08-7d818070e11f)]",
        ["source"] = new Dictionary<string, object>()
        {
        ["name"] = "str",
        ["kind"] = "Solution",
        ["sourceId"] = "str.azure-sentinel-solution-str"},
        ["support"] = new Dictionary<string, object>()
        {
        ["name"] = "Microsoft Corporation",
        ["email"] = "support@microsoft.com",
        ["link"] = "https://support.microsoft.com/",
        ["tier"] = "Microsoft"},
        ["version"] = "1.0.0"}} }
    }),
};
ArmOperation<SecurityInsightsTemplateResource> lro = await securityInsightsTemplate.UpdateAsync(WaitUntil.Completed, data);
SecurityInsightsTemplateResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsTemplateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
