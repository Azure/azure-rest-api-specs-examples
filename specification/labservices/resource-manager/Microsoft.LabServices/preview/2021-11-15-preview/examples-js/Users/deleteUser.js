const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to delete a user resource.
 *
 * @summary Operation to delete a user resource.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Users/deleteUser.json
 */
async function deleteUser() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labName = "testlab";
  const userName = "testuser";
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.users.beginDeleteAndWait(resourceGroupName, labName, userName);
  console.log(result);
}

deleteUser().catch(console.error);
