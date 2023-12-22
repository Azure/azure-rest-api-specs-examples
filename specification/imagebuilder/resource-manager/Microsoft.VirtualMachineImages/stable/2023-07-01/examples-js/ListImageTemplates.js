const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the VM image templates associated with the subscription.
 *
 * @summary Gets information about the VM image templates associated with the subscription.
 * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2023-07-01/examples/ListImageTemplates.json
 */
async function listImagesBySubscription() {
  const subscriptionId = process.env["IMAGEBUILDER_SUBSCRIPTION_ID"] || "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new ImageBuilderClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineImageTemplates.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
