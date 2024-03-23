const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts the SAP application, that is the Central Services instance and Application server instances.
 *
 * @summary Starts the SAP application, that is the Central Services instance and Application server instances.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Start.json
 */
async function startTheSapSystem() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "test-rg";
  const sapVirtualInstanceName = "X00";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPVirtualInstances.beginStartAndWait(
    resourceGroupName,
    sapVirtualInstanceName,
  );
  console.log(result);
}
