const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a custom recommendation over a given scope
 *
 * @summary delete a custom recommendation over a given scope
 * x-ms-original-file: 2024-08-01/CustomRecommendations/DeleteByManagementGroupCustomRecommendation_example.json
 */
async function deleteACustomRecommendationOverManagementGroupScope() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  await client.customRecommendations.delete(
    "providers/Microsoft.Management/managementGroups/contoso",
    "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
  );
}
