const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enable remote vendor management of the provided storage appliance.
 *
 * @summary Enable remote vendor management of the provided storage appliance.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/StorageAppliances_EnableRemoteVendorManagement.json
 */
async function turnOnRemoteVendorManagementForStorageAppliance() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const storageApplianceName = "storageApplianceName";
  const storageApplianceEnableRemoteVendorManagementParameters = {
    supportEndpoints: ["10.0.0.0/24"],
  };
  const options = {
    storageApplianceEnableRemoteVendorManagementParameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.storageAppliances.beginEnableRemoteVendorManagementAndWait(
    resourceGroupName,
    storageApplianceName,
    options
  );
  console.log(result);
}
