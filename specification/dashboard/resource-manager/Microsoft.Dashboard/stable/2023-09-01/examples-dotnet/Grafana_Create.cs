using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Grafana;
using Azure.ResourceManager.Grafana.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/Grafana_Create.json
// this example is just showing the usage of "Grafana_Create" operation, for the dependent resources, they will have to be created separately.

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
    Properties = new ManagedGrafanaProperties()
    {
        PublicNetworkAccess = GrafanaPublicNetworkAccess.Enabled,
        ZoneRedundancy = GrafanaZoneRedundancy.Enabled,
        ApiKey = GrafanaApiKey.Enabled,
        DeterministicOutboundIP = DeterministicOutboundIP.Enabled,
        MonitorWorkspaceIntegrations =
        {
        new MonitorWorkspaceIntegration()
        {
        MonitorWorkspaceResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.monitor/accounts/myAzureMonitorWorkspace"),
        }
        },
        EnterpriseConfigurations = new EnterpriseConfigurations()
        {
            MarketplacePlanId = "myPlanId",
            MarketplaceAutoRenew = MarketplaceAutoRenew.Enabled,
        },
        GrafanaConfigurationsSmtp = new Smtp()
        {
            Enabled = true,
            Host = "smtp.sendemail.com:587",
            User = "username",
            Password = "<password>",
            FromAddress = "test@sendemail.com",
            FromName = "emailsender",
            StartTLSPolicy = StartTLSPolicy.OpportunisticStartTLS,
            SkipVerify = true,
        },
        GrafanaPlugins =
        {
        ["sample-plugin-id"] = new GrafanaPlugin(),
        },
        GrafanaMajorVersion = "9",
    },
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Tags =
    {
    ["Environment"] = "Dev",
    },
};
ArmOperation<ManagedGrafanaResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, workspaceName, data);
ManagedGrafanaResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedGrafanaData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
