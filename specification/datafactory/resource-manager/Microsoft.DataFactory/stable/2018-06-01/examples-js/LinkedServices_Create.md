Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.3.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function linkedServicesCreate() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const linkedServiceName = "exampleLinkedService";
  const linkedService = {
    properties: {
      type: "AzureStorage",
      connectionString: {
        type: "SecureString",
        value:
          "DefaultEndpointsProtocol=https;AccountName=examplestorageaccount;AccountKey=<storage key>",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.linkedServices.createOrUpdate(
    resourceGroupName,
    factoryName,
    linkedServiceName,
    linkedService
  );
  console.log(result);
}

linkedServicesCreate().catch(console.error);
```
