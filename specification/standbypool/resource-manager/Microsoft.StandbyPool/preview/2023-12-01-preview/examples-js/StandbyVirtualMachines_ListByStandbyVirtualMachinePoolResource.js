const { StandbyPoolManagementClient } = require("@azure/arm-standbypool");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List StandbyVirtualMachineResource resources by StandbyVirtualMachinePoolResource
 *
 * @summary List StandbyVirtualMachineResource resources by StandbyVirtualMachinePoolResource
 * x-ms-original-file: specification/standbypool/resource-manager/Microsoft.StandbyPool/preview/2023-12-01-preview/examples/StandbyVirtualMachines_ListByStandbyVirtualMachinePoolResource.json
 */
async function standbyVirtualMachinesListByStandbyVirtualMachinePoolResource() {
  const subscriptionId =
    process.env["STANDBYPOOL_SUBSCRIPTION_ID"] || "8CC31D61-82D7-4B2B-B9DC-6B924DE7D229";
  const resourceGroupName = process.env["STANDBYPOOL_RESOURCE_GROUP"] || "rgstandbypool";
  const standbyVirtualMachinePoolName = "pool";
  const credential = new DefaultAzureCredential();
  const client = new StandbyPoolManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.standbyVirtualMachines.listByStandbyVirtualMachinePoolResource(
    resourceGroupName,
    standbyVirtualMachinePoolName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
