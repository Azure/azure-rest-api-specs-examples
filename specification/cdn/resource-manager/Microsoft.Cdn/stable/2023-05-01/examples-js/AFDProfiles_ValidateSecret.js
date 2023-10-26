const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validate a Secret in the profile.
 *
 * @summary Validate a Secret in the profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/AFDProfiles_ValidateSecret.json
 */
async function validateSecret() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const validateSecretInput = {
    secretSource: {
      id: "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/certificate/certName",
    },
    secretType: "CustomerCertificate",
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdProfiles.validateSecret(
    resourceGroupName,
    profileName,
    validateSecretInput
  );
  console.log(result);
}
