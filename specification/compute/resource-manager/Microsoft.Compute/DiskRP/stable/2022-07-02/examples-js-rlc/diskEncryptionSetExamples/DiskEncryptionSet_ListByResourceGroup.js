const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { paginate } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Lists all the disk encryption sets under a resource group.
 *
 * @summary Lists all the disk encryption sets under a resource group.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_ListByResourceGroup.json
 */
async function listAllDiskEncryptionSetsInAResourceGroup() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "myResourceGroup";
  const options = {
    queryParameters: { "api-version": "2022-07-02" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets",
      subscriptionId,
      resourceGroupName
    )
    .get(options);
  const pageData = paginate(client, initialResponse);
  const result = [];
  for await (const item of pageData) {
    result.push(item);
  }
  console.log(result);
}

listAllDiskEncryptionSetsInAResourceGroup().catch(console.error);
