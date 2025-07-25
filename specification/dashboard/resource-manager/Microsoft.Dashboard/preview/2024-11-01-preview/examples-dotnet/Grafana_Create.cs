using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Grafana.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Grafana;

// Generated from example definition: 2024-11-01-preview/Grafana_Create.json
// this example is just showing the usage of "ManagedGrafana_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ManagedGrafanaResource
ManagedGrafanaCollection collection = resourceGroupResource.GetManagedGrafanas();

// invoke the operation
string workspaceName = "myWorkspace";
ManagedGrafanaData data = new ManagedGrafanaData(new AzureLocation("West US"))
{
    SkuName = "Standard",
    Properties = new ManagedGrafanaProperties
    {
        PublicNetworkAccess = GrafanaPublicNetworkAccess.Enabled,
        ZoneRedundancy = GrafanaZoneRedundancy.Enabled,
        ApiKey = GrafanaApiKey.Enabled,
        DeterministicOutboundIP = DeterministicOutboundIP.Enabled,
        MonitorWorkspaceIntegrations = {new MonitorWorkspaceIntegration
        {
        MonitorWorkspaceResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.monitor/accounts/myAzureMonitorWorkspace"),
        }},
        EnterpriseConfigurations = new EnterpriseConfigurations
        {
            MarketplacePlanId = "myPlanId",
            MarketplaceAutoRenew = MarketplaceAutoRenew.Enabled,
        },
        GrafanaConfigurations = new GrafanaConfigurations
        {
            Smtp = new GrafanaSmtpSettings
            {
                IsEnabled = true,
                Host = "smtp.sendemail.com:587",
                User = "username",
                Password = "<password>",
                FromAddress = "test@sendemail.com",
                FromName = "emailsender",
                StartTLSPolicy = GrafanaStartTlsPolicy.OpportunisticStartTls,
                SkipVerify = true,
            },
            IsExternalEnabled = true,
            Users = new GrafanaUserSettings
            {
                ViewersCanEdit = true,
                EditorsCanAdmin = true,
            },
            IsCsrfAlwaysCheckEnabled = false,
            IsCaptureEnabled = false,
        },
        GrafanaPlugins =
        {
        ["sample-plugin-id"] = new GrafanaPlugin()
        },
        GrafanaMajorVersion = "9",
    },
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Tags =
    {
    ["Environment"] = "Dev"
    },
};
ArmOperation<ManagedGrafanaResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, workspaceName, data);
ManagedGrafanaResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedGrafanaData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
