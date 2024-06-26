const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the SAP Central Services Instance resource for the given Virtual Instance for SAP solutions resource.
 *
 * @summary Lists the SAP Central Services Instance resource for the given Virtual Instance for SAP solutions resource.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPCentralInstances_List.json
 */
async function sapCentralInstancesList() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "6d875e77-e412-4d7d-9af4-8895278b4443";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "test-rg";
  const sapVirtualInstanceName = "X00";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sAPCentralInstances.list(
    resourceGroupName,
    sapVirtualInstanceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
