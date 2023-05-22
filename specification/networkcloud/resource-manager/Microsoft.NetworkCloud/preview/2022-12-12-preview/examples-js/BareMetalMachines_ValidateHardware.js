const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validate the hardware of the provided bare metal machine.
 *
 * @summary Validate the hardware of the provided bare metal machine.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachines_ValidateHardware.json
 */
async function validateTheBareMetalMachineHardware() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const bareMetalMachineName = "bareMetalMachineName";
  const bareMetalMachineValidateHardwareParameters = {
    validationCategory: "BasicValidation",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.bareMetalMachines.beginValidateHardwareAndWait(
    resourceGroupName,
    bareMetalMachineName,
    bareMetalMachineValidateHardwareParameters
  );
  console.log(result);
}
