const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the security connectors in the specified resource group. Use the 'nextLink' property in the response to get the next page of security connectors for the specified resource group.
 *
 * @summary Lists all the security connectors in the specified resource group. Use the 'nextLink' property in the response to get the next page of security connectors for the specified resource group.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-05-01-preview/examples/SecurityConnectors/GetSecurityConnectorsResourceGroup_example.json
 */
async function listAllSecurityConnectorsOfASpecifiedResourceGroup() {
  const subscriptionId = "a5caac9c-5c04-49af-b3d0-e204f40345d5";
  const resourceGroupName = "exampleResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.securityConnectors.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllSecurityConnectorsOfASpecifiedResourceGroup().catch(console.error);
