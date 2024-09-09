using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/entityQueries/CreateEntityQueryActivity.json
// this example is just showing the usage of "EntityQueries_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this SecurityInsightsEntityQueryResource
SecurityInsightsEntityQueryCollection collection = operationalInsightsWorkspaceSecurityInsights.GetSecurityInsightsEntityQueries();

// invoke the operation
string entityQueryId = "07da3cc8-c8ad-4710-a44e-334cdcb7882b";
SecurityInsightsEntityQueryCreateOrUpdateContent content = new ActivityCustomEntityQuery()
{
    Title = "An account was deleted on this host",
    Content = "On '{{Computer}}' the account '{{TargetAccount}}' was deleted by '{{AddedBy}}'",
    Description = "Account deleted on host",
    Query = "let GetAccountActions = (v_Host_Name:string, v_Host_NTDomain:string, v_Host_DnsDomain:string, v_Host_AzureID:string, v_Host_OMSAgentID:string){\nSecurityEvent\n| where EventID in (4725, 4726, 4767, 4720, 4722, 4723, 4724)\n// parsing for Host to handle variety of conventions coming from data\n| extend Host_HostName = case(\nComputer has '@', tostring(split(Computer, '@')[0]),\nComputer has '\\\\', tostring(split(Computer, '\\\\')[1]),\nComputer has '.', tostring(split(Computer, '.')[0]),\nComputer\n)\n| extend Host_NTDomain = case(\nComputer has '\\\\', tostring(split(Computer, '\\\\')[0]), \nComputer has '.', tostring(split(Computer, '.')[-2]), \nComputer\n)\n| extend Host_DnsDomain = case(\nComputer has '\\\\', tostring(split(Computer, '\\\\')[0]), \nComputer has '.', strcat_array(array_slice(split(Computer,'.'),-2,-1),'.'), \nComputer\n)\n| where (Host_HostName =~ v_Host_Name and Host_NTDomain =~ v_Host_NTDomain) \nor (Host_HostName =~ v_Host_Name and Host_DnsDomain =~ v_Host_DnsDomain) \nor v_Host_AzureID =~ _ResourceId \nor v_Host_OMSAgentID == SourceComputerId\n| project TimeGenerated, EventID, Activity, Computer, TargetAccount, TargetUserName, TargetDomainName, TargetSid, SubjectUserName, SubjectUserSid, _ResourceId, SourceComputerId\n| extend AddedBy = SubjectUserName\n// Future support for Activities\n| extend timestamp = TimeGenerated, HostCustomEntity = Computer, AccountCustomEntity = TargetAccount\n};\nGetAccountActions('{{Host_HostName}}', '{{Host_NTDomain}}', '{{Host_DnsDomain}}', '{{Host_AzureID}}', '{{Host_OMSAgentID}}')\n \n| where EventID == 4726 ",
    InputEntityType = SecurityInsightsEntityType.Host,
    RequiredInputFieldsSets =
    {
    new string[]
    {
    "Host_HostName","Host_NTDomain"
    },new string[]
    {
    "Host_HostName","Host_DnsDomain"
    },new string[]
    {
    "Host_AzureID"
    },new string[]
    {
    "Host_OMSAgentID"
    }
    },
    EntitiesFilter =
    {
    ["Host_OsFamily"] = new string[]
    {
    "Windows"
    },
    },
    TemplateName = null,
    IsEnabled = true,
    ETag = new ETag("\"0300bf09-0000-0000-0000-5c37296e0000\""),
};
ArmOperation<SecurityInsightsEntityQueryResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, entityQueryId, content);
SecurityInsightsEntityQueryResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsEntityQueryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
