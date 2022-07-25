const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a Maps Creator resource.
 *
 * @summary Delete a Maps Creator resource.
 * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/DeleteMapsCreator.json
 */
async function deleteCreatorResource() {
  const subscriptionId = "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
  const resourceGroupName = "myResourceGroup";
  const accountName = "myMapsAccount";
  const creatorName = "myCreator";
  const credential = new DefaultAzureCredential();
  const client = new AzureMapsManagementClient(credential, subscriptionId);
  const result = await client.creators.delete(resourceGroupName, accountName, creatorName);
  console.log(result);
}

deleteCreatorResource().catch(console.error);
