using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/metadata/PutMetadata.json
// this example is just showing the usage of "Metadata_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OperationalInsightsWorkspaceSecurityInsightsResource created on azure
// for more information of creating OperationalInsightsWorkspaceSecurityInsightsResource, please refer to the document of OperationalInsightsWorkspaceSecurityInsightsResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
ResourceIdentifier operationalInsightsWorkspaceSecurityInsightsResourceId = OperationalInsightsWorkspaceSecurityInsightsResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
OperationalInsightsWorkspaceSecurityInsightsResource operationalInsightsWorkspaceSecurityInsights = client.GetOperationalInsightsWorkspaceSecurityInsightsResource(operationalInsightsWorkspaceSecurityInsightsResourceId);

// get the collection of this SecurityInsightsMetadataResource
SecurityInsightsMetadataCollection collection = operationalInsightsWorkspaceSecurityInsights.GetAllSecurityInsightsMetadata();

// invoke the operation
string metadataName = "metadataName";
SecurityInsightsMetadataData data = new SecurityInsightsMetadataData()
{
    ContentId = "c00ee137-7475-47c8-9cce-ec6f0f1bedd0",
    ParentId = "/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName",
    Version = "1.0.0.0",
    Kind = "AnalyticsRule",
    Source = new SecurityInsightsMetadataSource(SecurityInsightsSourceKind.Solution)
    {
        Name = "Contoso Solution 1.0",
        SourceId = "b688a130-76f4-4a07-bf57-762222a3cadf",
    },
    Author = new SecurityInsightsMetadataAuthor()
    {
        Name = "User Name",
        Email = "email@microsoft.com",
    },
    Support = new SecurityInsightsMetadataSupport(SecurityInsightsSupportTier.Partner)
    {
        Name = "Microsoft",
        Email = "support@microsoft.com",
        Link = "https://support.microsoft.com/",
    },
    Dependencies = new SecurityInsightsMetadataDependencies()
    {
        Operator = new ThreatIntelligenceQueryOperator("AND"),
        Criteria =
        {
        new SecurityInsightsMetadataDependencies()
        {
        Operator = new ThreatIntelligenceQueryOperator("OR"),
        Criteria =
        {
        new SecurityInsightsMetadataDependencies()
        {
        ContentId = "045d06d0-ee72-4794-aba4-cf5646e4c756",
        Kind = SecurityInsightsKind.DataConnector,
        Name = "Microsoft Defender for Endpoint",
        },new SecurityInsightsMetadataDependencies()
        {
        ContentId = "dbfcb2cc-d782-40ef-8d94-fe7af58a6f2d",
        Kind = SecurityInsightsKind.DataConnector,
        },new SecurityInsightsMetadataDependencies()
        {
        ContentId = "de4dca9b-eb37-47d6-a56f-b8b06b261593",
        Kind = SecurityInsightsKind.DataConnector,
        Version = "2.0",
        }
        },
        },new SecurityInsightsMetadataDependencies()
        {
        ContentId = "31ee11cc-9989-4de8-b176-5e0ef5c4dbab",
        Kind = SecurityInsightsKind.Playbook,
        Version = "1.0",
        },new SecurityInsightsMetadataDependencies()
        {
        ContentId = "21ba424a-9438-4444-953a-7059539a7a1b",
        Kind = SecurityInsightsKind.Parser,
        }
        },
    },
    Categories = new SecurityInsightsMetadataCategories()
    {
        Domains =
        {
        "Application","Security – Insider Threat"
        },
        Verticals =
        {
        "Healthcare"
        },
    },
    Providers =
    {
    "Amazon","Microsoft"
    },
    FirstPublishOn = DateTimeOffset.Parse("2021-05-18"),
    LastPublishOn = DateTimeOffset.Parse("2021-05-18"),
    CustomVersion = "1.0",
    ContentSchemaVersion = "2.0",
    ThreatAnalysisTactics =
    {
    "reconnaissance","commandandcontrol"
    },
    ThreatAnalysisTechniques =
    {
    "T1548","T1548.001"
    },
    PreviewImages =
    {
    "firstImage.png","secondImage.jpeg"
    },
    PreviewImagesDark =
    {
    "firstImageDark.png","secondImageDark.jpeg"
    },
};
ArmOperation<SecurityInsightsMetadataResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, metadataName, data);
SecurityInsightsMetadataResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsMetadataData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
