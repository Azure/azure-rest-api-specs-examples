const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all of the available operations the Desktop Virtualization resource provider supports.
 *
 * @summary List all of the available operations the Desktop Virtualization resource provider supports.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/OperationDescription_List.json
 */
async function operationDescriptionList() {
  const subscriptionId =
    process.env["DESKTOPVIRTUALIZATION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
