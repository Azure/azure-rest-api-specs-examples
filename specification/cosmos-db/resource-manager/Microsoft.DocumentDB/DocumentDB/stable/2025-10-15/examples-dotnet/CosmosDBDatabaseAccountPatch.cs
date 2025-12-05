using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBDatabaseAccountPatch.json
// this example is just showing the usage of "DatabaseAccounts_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBAccountResource created on azure
// for more information of creating CosmosDBAccountResource, please refer to the document of CosmosDBAccountResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
ResourceIdentifier cosmosDBAccountResourceId = CosmosDBAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
CosmosDBAccountResource cosmosDBAccount = client.GetCosmosDBAccountResource(cosmosDBAccountResourceId);

// invoke the operation
CosmosDBAccountPatch patch = new CosmosDBAccountPatch
{
    Tags =
    {
    ["dept"] = "finance"
    },
    Location = new AzureLocation("westus"),
    Identity = new ManagedServiceIdentity("SystemAssigned,UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1")] = new UserAssignedIdentity()
        },
    },
    ConsistencyPolicy = new ConsistencyPolicy(DefaultConsistencyLevel.BoundedStaleness)
    {
        MaxStalenessPrefix = 200L,
        MaxIntervalInSeconds = 10,
    },
    IPRules = {new CosmosDBIPAddressOrRange
    {
    IPAddressOrRange = "23.43.230.120",
    }, new CosmosDBIPAddressOrRange
    {
    IPAddressOrRange = "110.12.240.0/12",
    }},
    IsVirtualNetworkFilterEnabled = true,
    VirtualNetworkRules = {new CosmosDBVirtualNetworkRule
    {
    Id = new ResourceIdentifier("/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
    IgnoreMissingVnetServiceEndpoint = false,
    }},
    DefaultIdentity = "FirstPartyIdentity",
    IsFreeTierEnabled = false,
    IsAnalyticalStorageEnabled = true,
    AnalyticalStorageSchemaType = AnalyticalStorageSchemaType.WellDefined,
    BackupPolicy = new PeriodicModeBackupPolicy
    {
        PeriodicModeProperties = new PeriodicModeProperties
        {
            BackupIntervalInMinutes = 240,
            BackupRetentionIntervalInHours = 720,
            BackupStorageRedundancy = CosmosDBBackupStorageRedundancy.Local,
        },
    },
    NetworkAclBypass = NetworkAclBypass.AzureServices,
    NetworkAclBypassResourceIds = { new ResourceIdentifier("/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName") },
    CapacityTotalThroughputLimit = 2000,
    EnablePartitionMerge = true,
    MinimalTlsVersion = CosmosDBMinimalTlsVersion.Tls,
    EnableBurstCapacity = true,
    EnablePerRegionPerPartitionAutoscale = true,
    EnablePriorityBasedExecution = true,
    DefaultPriorityLevel = DefaultPriorityLevel.Low,
};
ArmOperation<CosmosDBAccountResource> lro = await cosmosDBAccount.UpdateAsync(WaitUntil.Completed, patch);
CosmosDBAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CosmosDBAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
