const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the security settings on a Data Box Edge/Data Box Gateway device.
 *
 * @summary Updates the security settings on a Data Box Edge/Data Box Gateway device.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/SecuritySettingsUpdatePost.json
 */
async function createOrUpdateSecuritySettings() {
  const subscriptionId =
    process.env["DATABOXEDGE_SUBSCRIPTION_ID"] || "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const resourceGroupName = process.env["DATABOXEDGE_RESOURCE_GROUP"] || "AzureVM";
  const securitySettings = {
    deviceAdminPassword: {
      encryptionAlgorithm: "AES256",
      encryptionCertThumbprint: "7DCBDFC44ED968D232C9A998FC105B5C70E84BE0",
      value: "<value>",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.devices.beginCreateOrUpdateSecuritySettingsAndWait(
    deviceName,
    resourceGroupName,
    securitySettings
  );
  console.log(result);
}
