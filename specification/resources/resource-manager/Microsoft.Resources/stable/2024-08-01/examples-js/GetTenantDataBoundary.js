const { DataboundariesManegementClient } = require("@azure/arm-databoundaries");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get data boundary of tenant.
 *
 * @summary Get data boundary of tenant.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-08-01/examples/GetTenantDataBoundary.json
 */
async function getDataBoundaryForTenant() {
  const defaultParam = "default";
  const credential = new DefaultAzureCredential();
  const client = new DataboundariesManegementClient(credential);
  const result = await client.dataBoundaries.getTenant(defaultParam);
  console.log(result);
}
