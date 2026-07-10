const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Azure Cosmos DB database account. The "Update" method is preferred when performing updates on an account.
 *
 * @summary creates or updates an Azure Cosmos DB database account. The "Update" method is preferred when performing updates on an account.
 * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountCreateMax.json
 */
async function cosmosDBDatabaseAccountCreateMax() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.databaseAccounts.createOrUpdate("rg1", "ddb1", {
    location: "westus",
    tags: {},
    kind: "MongoDB",
    identity: {
      type: "SystemAssigned,UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1":
          {},
      },
    },
    databaseAccountOfferType: "Standard",
    ipRules: [{ ipAddressOrRange: "23.43.230.120" }, { ipAddressOrRange: "110.12.240.0/12" }],
    isVirtualNetworkFilterEnabled: true,
    virtualNetworkRules: [
      {
        id: "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1",
        ignoreMissingVNetServiceEndpoint: false,
      },
    ],
    publicNetworkAccess: "Enabled",
    locations: [
      { failoverPriority: 0, locationName: "southcentralus", isZoneRedundant: false },
      { failoverPriority: 1, locationName: "eastus", isZoneRedundant: false },
    ],
    consistencyPolicy: {
      defaultConsistencyLevel: "BoundedStaleness",
      maxIntervalInSeconds: 10,
      maxStalenessPrefix: 200,
    },
    keyVaultKeyUri: "https://myKeyVault.vault.azure.net",
    defaultIdentity: "FirstPartyIdentity",
    enableFreeTier: false,
    apiProperties: { serverVersion: "3.2" },
    enableAnalyticalStorage: true,
    enableBurstCapacity: true,
    enablePriorityBasedExecution: true,
    defaultPriorityLevel: "Low",
    enablePerRegionPerPartitionAutoscale: true,
    analyticalStorageConfiguration: { schemaType: "WellDefined" },
    createMode: "Default",
    backupPolicy: {
      type: "Periodic",
      periodicModeProperties: {
        backupIntervalInMinutes: 240,
        backupRetentionIntervalInHours: 8,
        backupStorageRedundancy: "Geo",
      },
    },
    cors: [{ allowedOrigins: "https://test" }],
    networkAclBypass: "AzureServices",
    networkAclBypassResourceIds: [
      "/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName",
    ],
    capacity: { totalThroughputLimit: 2000 },
    enforceHierarchicalPartitionKeyIdLastLevel: false,
    minimalTlsVersion: "Tls12",
  });
  console.log(result);
}
