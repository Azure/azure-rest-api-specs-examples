const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of an existing Azure Cosmos DB database account.
 *
 * @summary Updates the properties of an existing Azure Cosmos DB database account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBDatabaseAccountPatch.json
 */
async function cosmosDbDatabaseAccountPatch() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const updateParameters = {
    analyticalStorageConfiguration: { schemaType: "WellDefined" },
    backupPolicy: {
      type: "Periodic",
      periodicModeProperties: {
        backupIntervalInMinutes: 240,
        backupRetentionIntervalInHours: 720,
        backupStorageRedundancy: "Geo",
      },
    },
    capacity: { totalThroughputLimit: 2000 },
    capacityMode: "Provisioned",
    consistencyPolicy: {
      defaultConsistencyLevel: "BoundedStaleness",
      maxIntervalInSeconds: 10,
      maxStalenessPrefix: 200,
    },
    defaultIdentity: "FirstPartyIdentity",
    defaultPriorityLevel: "Low",
    diagnosticLogSettings: { enableFullTextQuery: "True" },
    enableAnalyticalStorage: true,
    enableBurstCapacity: true,
    enableFreeTier: false,
    enablePartitionMerge: true,
    enablePerRegionPerPartitionAutoscale: true,
    enablePriorityBasedExecution: true,
    identity: {
      type: "SystemAssigned,UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/fa5fc227A624475eB696Cdd604c735bc/resourceGroups/eu2cgroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/id1":
          {},
      },
    },
    ipRules: [{ ipAddressOrRange: "23.43.230.120" }, { ipAddressOrRange: "110.12.240.0/12" }],
    isVirtualNetworkFilterEnabled: true,
    location: "westus",
    minimalTlsVersion: "Tls",
    networkAclBypass: "AzureServices",
    networkAclBypassResourceIds: [
      "/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName",
    ],
    tags: { dept: "finance" },
    virtualNetworkRules: [
      {
        id: "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1",
        ignoreMissingVNetServiceEndpoint: false,
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.databaseAccounts.beginUpdateAndWait(
    resourceGroupName,
    accountName,
    updateParameters,
  );
  console.log(result);
}
