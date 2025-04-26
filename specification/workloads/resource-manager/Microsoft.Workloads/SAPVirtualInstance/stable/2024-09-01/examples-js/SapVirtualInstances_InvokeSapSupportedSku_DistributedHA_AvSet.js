const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a list of SAP supported SKUs for ASCS, Application and Database tier.
 *
 * @summary get a list of SAP supported SKUs for ASCS, Application and Database tier.
 * x-ms-original-file: 2024-09-01/SapVirtualInstances_InvokeSapSupportedSku_DistributedHA_AvSet.json
 */
async function sapSupportedSKUsForDistributedHAEnvironmentWithAvailabilitySet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sapVirtualInstances.getSapSupportedSku("centralus", {
    appLocation: "eastus",
    sapProduct: "S4HANA",
    environment: "Prod",
    databaseType: "HANA",
    deploymentType: "ThreeTier",
    highAvailabilityType: "AvailabilitySet",
  });
  console.log(result);
}
