using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/alertRules/CreateFusionAlertRuleWithFusionScenarioExclusion.json
// this example is just showing the usage of "AlertRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsAlertRuleResource created on azure
// for more information of creating SecurityInsightsAlertRuleResource, please refer to the document of SecurityInsightsAlertRuleResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string ruleId = "myFirstFusionRule";
ResourceIdentifier securityInsightsAlertRuleResourceId = SecurityInsightsAlertRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, ruleId);
SecurityInsightsAlertRuleResource securityInsightsAlertRule = client.GetSecurityInsightsAlertRuleResource(securityInsightsAlertRuleResourceId);

// invoke the operation
SecurityInsightsAlertRuleData data = new SecurityInsightsFusionAlertRule()
{
    AlertRuleTemplateName = "f71aba3d-28fb-450b-b192-4e76a83015c8",
    IsEnabled = true,
    SourceSettings =
    {
    new FusionSourceSettings(true,"Anomalies")
    {
    SourceSubTypes =
    {
    },
    },new FusionSourceSettings(true,"Alert providers")
    {
    SourceSubTypes =
    {
    new FusionSourceSubTypeSetting(true,"Azure Active Directory Identity Protection",new FusionSubTypeSeverityFilter()
    {
    Filters =
    {
    new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.High,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Medium,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Low,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Informational,true)
    },
    }),new FusionSourceSubTypeSetting(true,"Azure Defender",new FusionSubTypeSeverityFilter()
    {
    Filters =
    {
    new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.High,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Medium,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Low,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Informational,true)
    },
    }),new FusionSourceSubTypeSetting(true,"Azure Defender for IoT",new FusionSubTypeSeverityFilter()
    {
    Filters =
    {
    new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.High,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Medium,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Low,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Informational,true)
    },
    }),new FusionSourceSubTypeSetting(true,"Microsoft 365 Defender",new FusionSubTypeSeverityFilter()
    {
    Filters =
    {
    new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.High,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Medium,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Low,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Informational,true)
    },
    }),new FusionSourceSubTypeSetting(true,"Microsoft Cloud App Security",new FusionSubTypeSeverityFilter()
    {
    Filters =
    {
    new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.High,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Medium,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Low,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Informational,true)
    },
    }),new FusionSourceSubTypeSetting(true,"Microsoft Defender for Endpoint",new FusionSubTypeSeverityFilter()
    {
    Filters =
    {
    new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.High,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Medium,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Low,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Informational,true)
    },
    }),new FusionSourceSubTypeSetting(true,"Microsoft Defender for Identity",new FusionSubTypeSeverityFilter()
    {
    Filters =
    {
    new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.High,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Medium,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Low,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Informational,true)
    },
    }),new FusionSourceSubTypeSetting(true,"Microsoft Defender for Office 365",new FusionSubTypeSeverityFilter()
    {
    Filters =
    {
    new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.High,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Medium,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Low,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Informational,true)
    },
    }),new FusionSourceSubTypeSetting(true,"Azure Sentinel scheduled analytics rules",new FusionSubTypeSeverityFilter()
    {
    Filters =
    {
    new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.High,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Medium,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Low,true),new FusionSubTypeSeverityFiltersItem(SecurityInsightsAlertSeverity.Informational,true)
    },
    })
    },
    },new FusionSourceSettings(true,"Raw logs from other sources")
    {
    SourceSubTypes =
    {
    new FusionSourceSubTypeSetting(true,"Palo Alto Networks",new FusionSubTypeSeverityFilter()
    {
    Filters =
    {
    },
    })
    },
    }
    },
    ETag = new ETag("3d00c3ca-0000-0100-0000-5d42d5010000"),
};
ArmOperation<SecurityInsightsAlertRuleResource> lro = await securityInsightsAlertRule.UpdateAsync(WaitUntil.Completed, data);
SecurityInsightsAlertRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsAlertRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
