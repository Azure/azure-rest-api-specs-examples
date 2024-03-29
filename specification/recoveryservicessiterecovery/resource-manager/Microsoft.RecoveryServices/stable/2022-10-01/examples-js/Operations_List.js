const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to return the list of available operations.
 *
 * @summary Operation to return the list of available operations.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/Operations_List.json
 */
async function returnsTheListOfAvailableOperations() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceGroupName = "resourceGroupPS1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

returnsTheListOfAvailableOperations().catch(console.error);
