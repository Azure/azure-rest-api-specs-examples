const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to use this method to get the list IoT Security solutions organized by resource group.
 *
 * @summary use this method to get the list IoT Security solutions organized by resource group.
 * x-ms-original-file: 2019-08-01/IoTSecuritySolutions/GetIoTSecuritySolutionsListByRg.json
 */
async function listIoTSecuritySolutionsByResourceGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.iotSecuritySolution.listByResourceGroup("MyGroup")) {
    resArray.push(item);
  }

  console.log(resArray);
}
