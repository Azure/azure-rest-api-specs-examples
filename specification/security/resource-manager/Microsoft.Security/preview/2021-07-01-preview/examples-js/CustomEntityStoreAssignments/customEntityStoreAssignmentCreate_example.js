const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a custom entity store assignment for the provided subscription, if not already exists.
 *
 * @summary Creates a custom entity store assignment for the provided subscription, if not already exists.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomEntityStoreAssignments/customEntityStoreAssignmentCreate_example.json
 */
async function createACustomEntityStoreAssignment() {
  const subscriptionId = "e5d1b86c-3051-44d5-8802-aa65d45a279b";
  const resourceGroupName = "TestResourceGroup";
  const customEntityStoreAssignmentName = "33e7cc6e-a139-4723-a0e5-76993aee0771";
  const customEntityStoreAssignmentRequestBody = {
    principal: "aaduser=f3923a3e-ad57-4752-b1a9-fbf3c8e5e082;72f988bf-86f1-41af-91ab-2d7cd011db47",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.customEntityStoreAssignments.create(
    resourceGroupName,
    customEntityStoreAssignmentName,
    customEntityStoreAssignmentRequestBody
  );
  console.log(result);
}

createACustomEntityStoreAssignment().catch(console.error);
