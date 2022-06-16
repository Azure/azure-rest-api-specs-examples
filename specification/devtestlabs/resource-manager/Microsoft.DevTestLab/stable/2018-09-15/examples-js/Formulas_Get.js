const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get formula.
 *
 * @summary Get formula.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Formulas_Get.json
 */
async function formulasGet() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{formulaName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.formulas.get(resourceGroupName, labName, name);
  console.log(result);
}

formulasGet().catch(console.error);
