const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a custom recommendation over a given scope
 *
 * @summary delete a custom recommendation over a given scope
 * x-ms-original-file: 2024-08-01/CustomRecommendations/DeleteBySecurityConnectorCustomRecommendation_example.json
 */
async function deleteACustomRecommendationOverSecurityConnectorScope() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  await client.customRecommendations.delete(
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector",
    "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
  );
}
