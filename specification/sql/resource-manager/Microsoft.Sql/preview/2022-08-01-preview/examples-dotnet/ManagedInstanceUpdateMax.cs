using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedInstanceUpdateMax.json
// this example is just showing the usage of "ManagedInstances_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceResource created on azure
// for more information of creating ManagedInstanceResource, please refer to the document of ManagedInstanceResource
string subscriptionId = "20D7082A-0FC7-4468-82BD-542694D5042B";
string resourceGroupName = "testrg";
string managedInstanceName = "testinstance";
ResourceIdentifier managedInstanceResourceId = ManagedInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName);
ManagedInstanceResource managedInstance = client.GetManagedInstanceResource(managedInstanceResourceId);

// invoke the operation
ManagedInstancePatch patch = new ManagedInstancePatch()
{
    Sku = new SqlSku("GP_Gen4")
    {
        Tier = "GeneralPurpose",
        Capacity = 8,
    },
    Tags =
    {
    ["tagKey1"] = "TagValue1",
    },
    AdministratorLogin = "dummylogin",
    AdministratorLoginPassword = "PLACEHOLDER",
    LicenseType = ManagedInstanceLicenseType.BasePrice,
    VCores = 8,
    StorageSizeInGB = 448,
    Collation = "SQL_Latin1_General_CP1_CI_AS",
    IsPublicDataEndpointEnabled = false,
    ProxyOverride = ManagedInstanceProxyOverride.Redirect,
    MaintenanceConfigurationId = new ResourceIdentifier("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1"),
    MinimalTlsVersion = "1.2",
    RequestedBackupStorageRedundancy = SqlBackupStorageRedundancy.Geo,
};
ArmOperation<ManagedInstanceResource> lro = await managedInstance.UpdateAsync(WaitUntil.Completed, patch);
ManagedInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
