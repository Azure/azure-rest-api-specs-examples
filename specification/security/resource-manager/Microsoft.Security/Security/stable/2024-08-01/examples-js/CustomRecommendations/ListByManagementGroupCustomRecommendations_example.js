const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a list of all relevant custom recommendations over a scope
 *
 * @summary get a list of all relevant custom recommendations over a scope
 * x-ms-original-file: 2024-08-01/CustomRecommendations/ListByManagementGroupCustomRecommendations_example.json
 */
async function listCustomRecommendationsByManagementGroupScope() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const resArray = new Array();
  for await (const item of client.customRecommendations.list(
    "providers/Microsoft.Management/managementGroups/contoso",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
