const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Use this method to delete yours IoT Security solution
 *
 * @summary Use this method to delete yours IoT Security solution
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/DeleteIoTSecuritySolution.json
 */
async function deleteAnIoTSecuritySolution() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = "MyGroup";
  const solutionName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.iotSecuritySolution.delete(resourceGroupName, solutionName);
  console.log(result);
}

deleteAnIoTSecuritySolution().catch(console.error);
