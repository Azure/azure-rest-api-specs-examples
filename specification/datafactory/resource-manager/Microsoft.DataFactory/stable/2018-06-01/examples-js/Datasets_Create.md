Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.4.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function datasetsCreate() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const datasetName = "exampleDataset";
  const dataset = {
    properties: {
      type: "AzureBlob",
      format: { type: "TextFormat" },
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

datasetsCreate().catch(console.error);
```
