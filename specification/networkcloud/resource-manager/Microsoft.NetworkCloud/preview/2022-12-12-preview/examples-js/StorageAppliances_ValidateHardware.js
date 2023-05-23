const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validate the hardware of the provided storage appliance.
 *
 * @summary Validate the hardware of the provided storage appliance.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/StorageAppliances_ValidateHardware.json
 */
async function validateTheStorageApplianceHardware() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const storageApplianceName = "storageApplianceName";
  const storageApplianceValidateHardwareParameters = {
    validationCategory: "BasicValidation",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.storageAppliances.beginValidateHardwareAndWait(
    resourceGroupName,
    storageApplianceName,
    storageApplianceValidateHardwareParameters
  );
  console.log(result);
}
