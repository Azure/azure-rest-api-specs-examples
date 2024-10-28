const { DataboundariesManegementClient } = require("@azure/arm-databoundaries");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Opt-in tenant to data boundary.
 *
 * @summary Opt-in tenant to data boundary.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-08-01/examples/PutDataBoundary.json
 */
async function optInToDataBoundary() {
  const defaultParam = "default";
  const dataBoundaryDefinition = {
    properties: { dataBoundary: "EU" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataboundariesManegementClient(credential);
  const result = await client.dataBoundaries.put(defaultParam, dataBoundaryDefinition);
  console.log(result);
}
