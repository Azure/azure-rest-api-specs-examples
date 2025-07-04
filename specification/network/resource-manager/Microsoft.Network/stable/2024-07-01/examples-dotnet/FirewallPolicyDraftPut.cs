using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/FirewallPolicyDraftPut.json
// this example is just showing the usage of "FirewallPolicyDrafts_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirewallPolicyDraftResource created on azure
// for more information of creating FirewallPolicyDraftResource, please refer to the document of FirewallPolicyDraftResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string firewallPolicyName = "firewallPolicy";
ResourceIdentifier firewallPolicyDraftResourceId = FirewallPolicyDraftResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, firewallPolicyName);
FirewallPolicyDraftResource firewallPolicyDraft = client.GetFirewallPolicyDraftResource(firewallPolicyDraftResourceId);

// invoke the operation
FirewallPolicyDraftData data = new FirewallPolicyDraftData
{
    ThreatIntelMode = AzureFirewallThreatIntelMode.Alert,
    ThreatIntelWhitelist = new FirewallPolicyThreatIntelWhitelist
    {
        IPAddresses = { "20.3.4.5" },
        Fqdns = { "*.microsoft.com" },
    },
    Insights = new FirewallPolicyInsights
    {
        IsEnabled = true,
        RetentionDays = 100,
        LogAnalyticsResources = new FirewallPolicyLogAnalyticsResources
        {
            Workspaces = {new FirewallPolicyLogAnalyticsWorkspace
            {
            Region = "westus",
            WorkspaceIdId = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace1"),
            }, new FirewallPolicyLogAnalyticsWorkspace
            {
            Region = "eastus",
            WorkspaceIdId = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace2"),
            }},
            DefaultWorkspaceIdId = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/defaultWorkspace"),
        },
    },
    Snat = new FirewallPolicySnat
    {
        PrivateRanges = { "IANAPrivateRanges" },
    },
    AllowSqlRedirect = true,
    DnsSettings = new DnsSettings
    {
        Servers = { "30.3.4.5" },
        EnableProxy = true,
        RequireProxyForNetworkRules = false,
    },
    ExplicitProxy = new FirewallPolicyExplicitProxy
    {
        EnableExplicitProxy = true,
        HttpPort = 8087,
        HttpsPort = 8087,
        EnablePacFile = true,
        PacFilePort = 8087,
        PacFile = "https://tinawstorage.file.core.windows.net/?sv=2020-02-10&ss=bfqt&srt=sco&sp=rwdlacuptfx&se=2021-06-04T07:01:12Z&st=2021-06-03T23:01:12Z&sip=68.65.171.11&spr=https&sig=Plsa0RRVpGbY0IETZZOT6znOHcSro71LLTTbzquYPgs%3D",
    },
    IntrusionDetection = new FirewallPolicyIntrusionDetection
    {
        Mode = FirewallPolicyIntrusionDetectionStateType.Alert,
        Profile = new FirewallPolicyIntrusionDetectionProfileType("Balanced"),
        Configuration = new FirewallPolicyIntrusionDetectionConfiguration
        {
            SignatureOverrides = {new FirewallPolicyIntrusionDetectionSignatureSpecification
            {
            Id = "2525004",
            Mode = FirewallPolicyIntrusionDetectionStateType.Deny,
            }},
            BypassTrafficSettings = {new FirewallPolicyIntrusionDetectionBypassTrafficSpecifications
            {
            Name = "bypassRule1",
            Description = "Rule 1",
            Protocol = FirewallPolicyIntrusionDetectionProtocol.Tcp,
            SourceAddresses = {"1.2.3.4"},
            DestinationAddresses = {"5.6.7.8"},
            DestinationPorts = {"*"},
            }},
        },
    },
};
ArmOperation<FirewallPolicyDraftResource> lro = await firewallPolicyDraft.CreateOrUpdateAsync(WaitUntil.Completed, data);
FirewallPolicyDraftResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FirewallPolicyDraftData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
