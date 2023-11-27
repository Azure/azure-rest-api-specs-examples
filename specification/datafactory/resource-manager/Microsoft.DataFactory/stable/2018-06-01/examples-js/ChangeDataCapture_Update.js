const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a change data capture resource.
 *
 * @summary Creates or updates a change data capture resource.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ChangeDataCapture_Update.json
 */
async function changeDataCaptureUpdate() {
  const subscriptionId =
    process.env["DATAFACTORY_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = process.env["DATAFACTORY_RESOURCE_GROUP"] || "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const changeDataCaptureName = "exampleChangeDataCapture";
  const changeDataCapture = {
    description:
      "Sample demo change data capture to transfer data from delimited (csv) to Azure SQL Database. Updating table mappings.",
    allowVNetOverride: false,
    status: "Stopped",
    sourceConnectionsInfo: [],
    targetConnectionsInfo: [],
    policy: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.changeDataCapture.createOrUpdate(
    resourceGroupName,
    factoryName,
    changeDataCaptureName,
    changeDataCapture
  );
  console.log(result);
}
