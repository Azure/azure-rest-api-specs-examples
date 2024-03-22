const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the SAP Application Server Instance resources for a given Virtual Instance for SAP solutions resource.
 *
 * @summary Lists the SAP Application Server Instance resources for a given Virtual Instance for SAP solutions resource.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapapplicationinstances/SAPApplicationServerInstances_List.json
 */
async function sapApplicationServerInstancesList() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "6d875e77-e412-4d7d-9af4-8895278b4443";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "test-rg";
  const sapVirtualInstanceName = "X00";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sAPApplicationServerInstances.list(
    resourceGroupName,
    sapVirtualInstanceName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
