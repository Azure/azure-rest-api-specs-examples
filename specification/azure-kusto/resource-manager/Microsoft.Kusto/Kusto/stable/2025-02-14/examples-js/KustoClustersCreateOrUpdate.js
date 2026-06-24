const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Kusto cluster.
 *
 * @summary create or update a Kusto cluster.
 * x-ms-original-file: 2025-02-14/KustoClustersCreateOrUpdate.json
 */
async function kustoClustersCreateOrUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.createOrUpdate("kustorptest", "kustoCluster", {
    identity: { type: "SystemAssigned" },
    location: "westus",
    allowedIpRangeList: ["0.0.0.0/0"],
    enableAutoStop: true,
    enableDoubleEncryption: false,
    enablePurge: true,
    enableStreamingIngest: true,
    languageExtensions: {
      value: [
        { languageExtensionImageName: "Python3_10_8", languageExtensionName: "PYTHON" },
        { languageExtensionImageName: "R", languageExtensionName: "R" },
      ],
    },
    publicIPType: "DualStack",
    publicNetworkAccess: "Enabled",
    sku: { name: "Standard_L16as_v3", capacity: 2, tier: "Standard" },
  });
  console.log(result);
}
