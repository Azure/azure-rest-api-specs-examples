using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2021-02-01/examples/FirewallPolicyPut.json
// this example is just showing the usage of "FirewallPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirewallPolicyResource created on azure
// for more information of creating FirewallPolicyResource, please refer to the document of FirewallPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string firewallPolicyName = "firewallPolicy";
ResourceIdentifier firewallPolicyResourceId = FirewallPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, firewallPolicyName);
FirewallPolicyResource firewallPolicy = client.GetFirewallPolicyResource(firewallPolicyResourceId);

// invoke the operation
FirewallPolicyData data = new FirewallPolicyData()
{
    ThreatIntelMode = AzureFirewallThreatIntelMode.Alert,
    ThreatIntelWhitelist = new FirewallPolicyThreatIntelWhitelist()
    {
        IPAddresses =
        {
        "20.3.4.5"
        },
        Fqdns =
        {
        "*.microsoft.com"
        },
    },
    Insights = new FirewallPolicyInsights()
    {
        IsEnabled = true,
        RetentionDays = 100,
        LogAnalyticsResources = new FirewallPolicyLogAnalyticsResources()
        {
            Workspaces =
            {
            new FirewallPolicyLogAnalyticsWorkspace()
            {
            Region = "westus",
            WorkspaceIdId = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace1"),
            },new FirewallPolicyLogAnalyticsWorkspace()
            {
            Region = "eastus",
            WorkspaceIdId = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace2"),
            }
            },
            DefaultWorkspaceIdId = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/defaultWorkspace"),
        },
    },
    SnatPrivateRanges =
    {
    "IANAPrivateRanges"
    },
    DnsSettings = new DnsSettings()
    {
        Servers =
        {
        "30.3.4.5"
        },
        EnableProxy = true,
        RequireProxyForNetworkRules = false,
    },
    IntrusionDetection = new FirewallPolicyIntrusionDetection()
    {
        Mode = FirewallPolicyIntrusionDetectionStateType.Alert,
        Configuration = new FirewallPolicyIntrusionDetectionConfiguration()
        {
            SignatureOverrides =
            {
            new FirewallPolicyIntrusionDetectionSignatureSpecification()
            {
            Id = "2525004",
            Mode = FirewallPolicyIntrusionDetectionStateType.Deny,
            }
            },
            BypassTrafficSettings =
            {
            new FirewallPolicyIntrusionDetectionBypassTrafficSpecifications()
            {
            Name = "bypassRule1",
            Description = "Rule 1",
            Protocol = FirewallPolicyIntrusionDetectionProtocol.Tcp,
            SourceAddresses =
            {
            "1.2.3.4"
            },
            DestinationAddresses =
            {
            "5.6.7.8"
            },
            DestinationPorts =
            {
            "*"
            },
            }
            },
        },
    },
    TransportSecurityCertificateAuthority = new FirewallPolicyCertificateAuthority()
    {
        KeyVaultSecretId = "https://kv/secret",
        Name = "clientcert",
    },
    SkuTier = FirewallPolicySkuTier.Premium,
    Location = new AzureLocation("West US"),
    Tags =
    {
    ["key1"] = "value1",
    },
};
ArmOperation<FirewallPolicyResource> lro = await firewallPolicy.UpdateAsync(WaitUntil.Completed, data);
FirewallPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FirewallPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
