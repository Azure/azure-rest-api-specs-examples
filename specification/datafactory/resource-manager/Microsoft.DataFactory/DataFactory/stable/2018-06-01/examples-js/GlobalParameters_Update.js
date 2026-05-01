const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a Global parameter
 *
 * @summary creates or updates a Global parameter
 * x-ms-original-file: 2018-06-01/GlobalParameters_Update.json
 */
async function globalParametersUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.globalParameters.createOrUpdate(
    "exampleResourceGroup",
    "exampleFactoryName",
    "default",
    { properties: { waitTime: { type: "Int", value: 5 } } },
  );
  console.log(result);
}
