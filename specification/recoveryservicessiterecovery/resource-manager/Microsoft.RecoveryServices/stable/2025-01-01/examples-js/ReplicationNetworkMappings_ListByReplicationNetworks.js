const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all ASR network mappings for the specified network.
 *
 * @summary Lists all ASR network mappings for the specified network.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationNetworkMappings_ListByReplicationNetworks.json
 */
async function getsAllTheNetworkMappingsUnderANetwork() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "9112a37f-0f3e-46ec-9c00-060c6edca071";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "srcBvte2a14C27";
  const resourceName = "srce2avaultbvtaC27";
  const fabricName = "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac";
  const networkName = "e2267b5c-2650-49bd-ab3f-d66aae694c06";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.replicationNetworkMappings.listByReplicationNetworks(
    resourceGroupName,
    resourceName,
    fabricName,
    networkName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
