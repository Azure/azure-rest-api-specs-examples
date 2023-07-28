const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of consoles for the provided virtual machine.
 *
 * @summary Get a list of consoles for the provided virtual machine.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/Consoles_ListByVirtualMachine.json
 */
async function listConsolesOfTheVirtualMachine() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const virtualMachineName = "virtualMachineName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.consoles.listByVirtualMachine(
    resourceGroupName,
    virtualMachineName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
