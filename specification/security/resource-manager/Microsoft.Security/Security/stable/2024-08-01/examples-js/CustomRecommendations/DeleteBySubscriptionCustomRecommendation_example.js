const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a custom recommendation over a given scope
 *
 * @summary delete a custom recommendation over a given scope
 * x-ms-original-file: 2024-08-01/CustomRecommendations/DeleteBySubscriptionCustomRecommendation_example.json
 */
async function deleteACustomRecommendationOverSubscriptionScope() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  await client.customRecommendations.delete(
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
    "1f3afdf9-d0c9-4c3d-847f-89da613e70a8",
  );
}
