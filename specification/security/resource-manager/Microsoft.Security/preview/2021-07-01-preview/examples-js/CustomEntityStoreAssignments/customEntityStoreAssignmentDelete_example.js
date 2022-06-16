const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a custom entity store assignment by name for a provided subscription
 *
 * @summary Delete a custom entity store assignment by name for a provided subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomEntityStoreAssignments/customEntityStoreAssignmentDelete_example.json
 */
async function deleteACustomEntityStoreAssignment() {
  const subscriptionId = "e5d1b86c-3051-44d5-8802-aa65d45a279b";
  const resourceGroupName = "TestResourceGroup";
  const customEntityStoreAssignmentName = "33e7cc6e-a139-4723-a0e5-76993aee0771";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.customEntityStoreAssignments.delete(
    resourceGroupName,
    customEntityStoreAssignmentName
  );
  console.log(result);
}

deleteACustomEntityStoreAssignment().catch(console.error);
