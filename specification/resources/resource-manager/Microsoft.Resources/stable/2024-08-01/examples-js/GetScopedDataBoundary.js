const { DataboundariesManegementClient } = require("@azure/arm-databoundaries");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get data boundary at specified scope
 *
 * @summary Get data boundary at specified scope
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-08-01/examples/GetScopedDataBoundary.json
 */
async function getDataBoundaryAtScope() {
  const scope =
    "subscriptions/11111111-1111-1111-1111-111111111111/resourcegroups/my-resource-group";
  const defaultParam = "default";
  const credential = new DefaultAzureCredential();
  const client = new DataboundariesManegementClient(credential);
  const result = await client.dataBoundaries.getScope(scope, defaultParam);
  console.log(result);
}
