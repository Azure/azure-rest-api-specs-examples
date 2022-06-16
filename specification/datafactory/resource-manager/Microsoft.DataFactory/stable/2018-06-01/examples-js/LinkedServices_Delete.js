const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a linked service.
 *
 * @summary Deletes a linked service.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/LinkedServices_Delete.json
 */
async function linkedServicesDelete() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const linkedServiceName = "exampleLinkedService";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.linkedServices.delete(
    resourceGroupName,
    factoryName,
    linkedServiceName
  );
  console.log(result);
}

linkedServicesDelete().catch(console.error);
