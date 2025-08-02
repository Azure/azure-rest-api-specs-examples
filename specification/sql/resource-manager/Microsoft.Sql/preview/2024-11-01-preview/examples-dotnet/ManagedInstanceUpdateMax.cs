using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-11-01-preview/examples/ManagedInstanceUpdateMax.json
// this example is just showing the usage of "ManagedInstances_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceResource created on azure
// for more information of creating ManagedInstanceResource, please refer to the document of ManagedInstanceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string managedInstanceName = "testinstance";
ResourceIdentifier managedInstanceResourceId = ManagedInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName);
ManagedInstanceResource managedInstance = client.GetManagedInstanceResource(managedInstanceResourceId);

// invoke the operation
ManagedInstancePatch patch = new ManagedInstancePatch
{
    Sku = new SqlSku("GP_Gen5")
    {
        Tier = "GeneralPurpose",
        Capacity = 8,
    },
    Tags =
    {
    ["tagKey1"] = "TagValue1"
    },
    AdministratorLogin = "dummylogin",
    AdministratorLoginPassword = "PLACEHOLDER",
    LicenseType = ManagedInstanceLicenseType.BasePrice,
    HybridSecondaryUsage = HybridSecondaryUsage.Passive,
    VCores = 8,
    StorageSizeInGB = 448,
    Collation = "SQL_Latin1_General_CP1_CI_AS",
    IsPublicDataEndpointEnabled = false,
    ProxyOverride = ManagedInstanceProxyOverride.Redirect,
    MaintenanceConfigurationId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1"),
    MinimalTlsVersion = "1.2",
    RequestedBackupStorageRedundancy = SqlBackupStorageRedundancy.Geo,
    AuthenticationMetadata = AuthMetadataLookupMode.Windows,
    DatabaseFormat = ManagedInstanceDatabaseFormat.AlwaysUpToDate,
    RequestedLogicalAvailabilityZone = SqlAvailabilityZoneType.One,
};
ArmOperation<ManagedInstanceResource> lro = await managedInstance.UpdateAsync(WaitUntil.Completed, patch);
ManagedInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
