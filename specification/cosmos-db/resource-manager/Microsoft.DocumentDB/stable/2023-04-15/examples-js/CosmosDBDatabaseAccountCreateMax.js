const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Azure Cosmos DB database account. The "Update" method is preferred when performing updates on an account.
 *
 * @summary Creates or updates an Azure Cosmos DB database account. The "Update" method is preferred when performing updates on an account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/examples/CosmosDBDatabaseAccountCreateMax.json
 */
async function cosmosDbDatabaseAccountCreateMax() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const createUpdateParameters = {
    analyticalStorageConfiguration: { schemaType: "WellDefined" },
    apiProperties: { serverVersion: "3.2" },
    backupPolicy: {
      type: "Periodic",
      periodicModeProperties: {
        backupIntervalInMinutes: 240,
        backupRetentionIntervalInHours: 8,
        backupStorageRedundancy: "Geo",
      },
    },
    capacity: { totalThroughputLimit: 2000 },
    consistencyPolicy: {
      defaultConsistencyLevel: "BoundedStaleness",
      maxIntervalInSeconds: 10,
      maxStalenessPrefix: 200,
    },
    cors: [{ allowedOrigins: "https://test" }],
    createMode: "Default",
    databaseAccountOfferType: "Standard",
    defaultIdentity: "FirstPartyIdentity",
    enableAnalyticalStorage: true,
    enableFreeTier: false,
    identity: {
      type: "SystemAssigned,UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/fa5fc227A624475eB696Cdd604c735bc/resourceGroups/eu2cgroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/id1":
          {},
      },
    },
    ipRules: [{ ipAddressOrRange: "23.43.230.120" }, { ipAddressOrRange: "110.12.240.0/12" }],
    isVirtualNetworkFilterEnabled: true,
    keyVaultKeyUri: "https://myKeyVault.vault.azure.net",
    kind: "MongoDB",
    location: "westus",
    locations: [
      {
        failoverPriority: 0,
        isZoneRedundant: false,
        locationName: "southcentralus",
      },
      { failoverPriority: 1, isZoneRedundant: false, locationName: "eastus" },
    ],
    minimalTlsVersion: "Tls12",
    networkAclBypass: "AzureServices",
    networkAclBypassResourceIds: [
      "/subscriptions/subId/resourcegroups/rgName/providers/Microsoft.Synapse/workspaces/workspaceName",
    ],
    publicNetworkAccess: "Enabled",
    tags: {},
    virtualNetworkRules: [
      {
        id: "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1",
        ignoreMissingVNetServiceEndpoint: false,
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.databaseAccounts.beginCreateOrUpdateAndWait(
    resourceGroupName,
    accountName,
    createUpdateParameters
  );
  console.log(result);
}
