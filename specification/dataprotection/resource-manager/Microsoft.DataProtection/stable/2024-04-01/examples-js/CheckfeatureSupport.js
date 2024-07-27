const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validates if a feature is supported
 *
 * @summary Validates if a feature is supported
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/CheckfeatureSupport.json
 */
async function checkAzureVMBackupFeatureSupport() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "0b352192-dcac-4cc7-992e-a96190ccc68c";
  const location = "WestUS";
  const parameters = {
    featureType: "DataSourceType",
    objectType: "FeatureValidationRequest",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.dataProtection.checkFeatureSupport(location, parameters);
  console.log(result);
}
