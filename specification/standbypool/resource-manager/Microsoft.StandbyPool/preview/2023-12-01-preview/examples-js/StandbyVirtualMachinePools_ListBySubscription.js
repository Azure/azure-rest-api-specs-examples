const { StandbyPoolManagementClient } = require("@azure/arm-standbypool");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List StandbyVirtualMachinePoolResource resources by subscription ID
 *
 * @summary List StandbyVirtualMachinePoolResource resources by subscription ID
 * x-ms-original-file: specification/standbypool/resource-manager/Microsoft.StandbyPool/preview/2023-12-01-preview/examples/StandbyVirtualMachinePools_ListBySubscription.json
 */
async function standbyVirtualMachinePoolsListBySubscription() {
  const subscriptionId =
    process.env["STANDBYPOOL_SUBSCRIPTION_ID"] || "8CC31D61-82D7-4B2B-B9DC-6B924DE7D229";
  const credential = new DefaultAzureCredential();
  const client = new StandbyPoolManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.standbyVirtualMachinePools.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
