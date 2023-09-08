using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.RecoveryServices;
using Azure.ResourceManager.RecoveryServices.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/PUTVault_WithMonitoringSettings.json
// this example is just showing the usage of "Vaults_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "77777777-b0c6-47a2-b37c-d8e65a629c18";
string resourceGroupName = "Default-RecoveryServices-ResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this RecoveryServicesVaultResource
RecoveryServicesVaultCollection collection = resourceGroupResource.GetRecoveryServicesVaults();

// invoke the operation
string vaultName = "swaggerExample";
RecoveryServicesVaultData data = new RecoveryServicesVaultData(new AzureLocation("West US"))
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Properties = new RecoveryServicesVaultProperties()
    {
        PublicNetworkAccess = VaultPublicNetworkAccess.Enabled,
        MonitoringSettings = new VaultMonitoringSettings()
        {
            AzureMonitorAlertAlertsForAllJobFailures = RecoveryServicesAlertsState.Enabled,
            ClassicAlertAlertsForCriticalOperations = RecoveryServicesAlertsState.Disabled,
        },
    },
    Sku = new RecoveryServicesSku(RecoveryServicesSkuName.Standard),
};
ArmOperation<RecoveryServicesVaultResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, vaultName, data);
RecoveryServicesVaultResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RecoveryServicesVaultData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
