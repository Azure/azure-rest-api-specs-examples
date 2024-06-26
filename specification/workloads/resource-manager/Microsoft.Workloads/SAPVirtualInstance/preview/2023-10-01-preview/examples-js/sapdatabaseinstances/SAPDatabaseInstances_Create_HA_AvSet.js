const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates the Database resource corresponding to the Virtual Instance for SAP solutions resource. <br><br>This will be used by service only. PUT by end user will return a Bad Request error.
 *
 * @summary Creates the Database resource corresponding to the Virtual Instance for SAP solutions resource. <br><br>This will be used by service only. PUT by end user will return a Bad Request error.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapdatabaseinstances/SAPDatabaseInstances_Create_HA_AvSet.json
 */
async function createSapDatabaseInstancesForHaSystemWithAvailabilitySet() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "6d875e77-e412-4d7d-9af4-8895278b4443";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "test-rg";
  const sapVirtualInstanceName = "X00";
  const databaseInstanceName = "databaseServer";
  const body = {
    location: "eastus",
    properties: {},
    tags: {},
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPDatabaseInstances.beginCreateAndWait(
    resourceGroupName,
    sapVirtualInstanceName,
    databaseInstanceName,
    options,
  );
  console.log(result);
}
