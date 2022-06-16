const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete an application control machine group
 *
 * @summary Delete an application control machine group
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ApplicationWhitelistings/DeleteAdaptiveApplicationControls_example.json
 */
async function deleteAnApplicationControlMachineGroup() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ascLocation = "centralus";
  const groupName = "GROUP1";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.adaptiveApplicationControls.delete(ascLocation, groupName);
  console.log(result);
}

deleteAnApplicationControlMachineGroup().catch(console.error);
