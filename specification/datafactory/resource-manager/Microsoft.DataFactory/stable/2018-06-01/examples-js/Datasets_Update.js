const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a dataset.
 *
 * @summary Creates or updates a dataset.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Update.json
 */
async function datasetsUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const datasetName = "exampleDataset";
  const dataset = {
    properties: {
      type: "AzureBlob",
      format: { type: "TextFormat" },
      description: "Example description",
      fileName: { type: "Expression", value: "@dataset().MyFileName" },
      folderPath: { type: "Expression", value: "@dataset().MyFolderPath" },
      linkedServiceName: {
        type: "LinkedServiceReference",
        referenceName: "exampleLinkedService",
      },
      parameters: {
        myFileName: { type: "String" },
        myFolderPath: { type: "String" },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.datasets.createOrUpdate(
    resourceGroupName,
    factoryName,
    datasetName,
    dataset
  );
  console.log(result);
}

datasetsUpdate().catch(console.error);
