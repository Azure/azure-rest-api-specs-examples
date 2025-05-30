const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the SAP Application Server Instance corresponding to the Virtual Instance for SAP solutions resource.
 *
 * @summary gets the SAP Application Server Instance corresponding to the Virtual Instance for SAP solutions resource.
 * x-ms-original-file: 2024-09-01/SapApplicationServerInstances_Get.json
 */
async function sapApplicationServerInstancesGet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6d875e77-e412-4d7d-9af4-8895278b4443";
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sapApplicationServerInstances.get("test-rg", "X00", "app01");
  console.log(result);
}
