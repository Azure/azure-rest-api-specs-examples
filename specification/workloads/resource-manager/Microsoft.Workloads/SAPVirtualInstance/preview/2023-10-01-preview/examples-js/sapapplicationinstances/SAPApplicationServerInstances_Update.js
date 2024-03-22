const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the SAP Application server instance resource. This can be used to update tags on the resource.
 *
 * @summary Updates the SAP Application server instance resource. This can be used to update tags on the resource.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapapplicationinstances/SAPApplicationServerInstances_Update.json
 */
async function sapApplicationServerInstancesUpdate() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "6d875e77-e412-4d7d-9af4-8895278b4443";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "test-rg";
  const sapVirtualInstanceName = "X00";
  const applicationInstanceName = "app01";
  const body = {
    tags: { tag1: "value1" },
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPApplicationServerInstances.update(
    resourceGroupName,
    sapVirtualInstanceName,
    applicationInstanceName,
    options,
  );
  console.log(result);
}
