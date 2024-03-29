const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Content Key Policy including secret values
 *
 * @summary Get a Content Key Policy including secret values
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-get-with-secrets.json
 */
async function getAnContentKeyPolicyWithSecrets() {
  const subscriptionId =
    process.env["MEDIASERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MEDIASERVICES_RESOURCE_GROUP"] || "contoso";
  const accountName = "contosomedia";
  const contentKeyPolicyName = "PolicyWithMultipleOptions";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.contentKeyPolicies.getPolicyPropertiesWithSecrets(
    resourceGroupName,
    accountName,
    contentKeyPolicyName
  );
  console.log(result);
}
