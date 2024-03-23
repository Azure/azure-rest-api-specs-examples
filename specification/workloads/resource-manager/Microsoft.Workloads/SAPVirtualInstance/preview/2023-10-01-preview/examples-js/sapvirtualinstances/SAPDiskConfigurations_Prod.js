const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the SAP Disk Configuration Layout prod/non-prod SAP System.
 *
 * @summary Get the SAP Disk Configuration Layout prod/non-prod SAP System.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPDiskConfigurations_Prod.json
 */
async function sapDiskConfigurationsProd() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const location = "centralus";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPDiskConfigurations(location);
  console.log(result);
}
