const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validate a Secret in the profile.
 *
 * @summary Validate a Secret in the profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Validate_Secret.json
 */
async function validateSecret() {
  const subscriptionId = "subid";
  const validateSecretInput = {
    secretSource: {
      id: "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/certificate/certName",
    },
    secretType: "CustomerCertificate",
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.validate.secret(validateSecretInput);
  console.log(result);
}

validateSecret().catch(console.error);
