const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves the data policy manifest with the given policy mode.
 *
 * @summary This operation retrieves the data policy manifest with the given policy mode.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-09-01/examples/getDataPolicyManifest.json
 */
async function retrieveADataPolicyManifestByPolicyMode() {
  const policyMode = "Microsoft.KeyVault.Data";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential);
  const result = await client.dataPolicyManifests.getByPolicyMode(policyMode);
  console.log(result);
}
