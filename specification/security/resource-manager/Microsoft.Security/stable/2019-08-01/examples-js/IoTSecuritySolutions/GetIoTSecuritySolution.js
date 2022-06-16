const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to User this method to get details of a specific IoT Security solution based on solution name
 *
 * @summary User this method to get details of a specific IoT Security solution based on solution name
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolution.json
 */
async function getAIoTSecuritySolution() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = "MyGroup";
  const solutionName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.iotSecuritySolution.get(resourceGroupName, solutionName);
  console.log(result);
}

getAIoTSecuritySolution().catch(console.error);
