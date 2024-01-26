const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a linked service.
 *
 * @summary Creates or updates a linked service.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/LinkedServices_Create.json
 */
async function linkedServicesCreate() {
  const subscriptionId =
    process.env["DATAFACTORY_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = process.env["DATAFACTORY_RESOURCE_GROUP"] || "exampleResourceGroup";
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
    linkedService,
  );
  console.log(result);
}
