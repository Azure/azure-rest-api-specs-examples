const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the networks available in a vault.
 *
 * @summary Lists the networks available in a vault.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationNetworks_List.json
 */
async function getsTheListOfNetworksViewOnlyApi() {
  const subscriptionId = "9112a37f-0f3e-46ec-9c00-060c6edca071";
  const resourceName = "srce2avaultbvtaC27";
  const resourceGroupName = "srcBvte2a14C27";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.replicationNetworks.list(resourceName, resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsTheListOfNetworksViewOnlyApi().catch(console.error);
