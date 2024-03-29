const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the token that is used for correlating ingested telemetry with the resources in the subscription.
 *
 * @summary Returns the token that is used for correlating ingested telemetry with the resources in the subscription.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/ListTokensIngestionSetting_example.json
 */
async function listIngestionSettingTokens() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ingestionSettingName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.ingestionSettings.listTokens(ingestionSettingName);
  console.log(result);
}
