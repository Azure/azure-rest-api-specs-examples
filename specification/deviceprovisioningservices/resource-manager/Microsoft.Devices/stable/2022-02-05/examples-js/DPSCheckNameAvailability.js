const { IotDpsClient } = require("@azure/arm-deviceprovisioningservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check if a provisioning service name is available. This will validate if the name is syntactically valid and if the name is usable
 *
 * @summary Check if a provisioning service name is available. This will validate if the name is syntactically valid and if the name is usable
 * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSCheckNameAvailability.json
 */
async function dpsCheckName() {
  const subscriptionId =
    process.env["DEVICEPROVISIONINGSERVICES_SUBSCRIPTION_ID"] ||
    "91d12660-3dec-467a-be2a-213b5544ddc0";
  const argumentsParam = { name: "test213123" };
  const credential = new DefaultAzureCredential();
  const client = new IotDpsClient(credential, subscriptionId);
  const result = await client.iotDpsResource.checkProvisioningServiceNameAvailability(
    argumentsParam
  );
  console.log(result);
}
