const { QumuloStorage } = require("@azure/arm-liftrqumulo");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a FileSystemResource
 *
 * @summary Update a FileSystemResource
 * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_Update_MinimumSet_Gen.json
 */
async function fileSystemsUpdateMinimumSetGen() {
  const subscriptionId = process.env["LIFTRQUMULO_SUBSCRIPTION_ID"] || "aaaaaaa";
  const resourceGroupName = process.env["LIFTRQUMULO_RESOURCE_GROUP"] || "rgQumulo";
  const fileSystemName = "aaaaaaaaaaaaaaaaa";
  const properties = {};
  const credential = new DefaultAzureCredential();
  const client = new QumuloStorage(credential, subscriptionId);
  const result = await client.fileSystems.update(resourceGroupName, fileSystemName, properties);
  console.log(result);
}
