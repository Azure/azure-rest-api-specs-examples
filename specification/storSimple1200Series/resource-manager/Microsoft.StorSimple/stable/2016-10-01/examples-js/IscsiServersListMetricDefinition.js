const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves metric definitions for all metrics aggregated at iSCSI server.
 *
 * @summary Retrieves metric definitions for all metrics aggregated at iSCSI server.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiServersListMetricDefinition.json
 */
async function iscsiServersListMetricDefinition() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-UGU4PITWNI";
  const iscsiServerName = "HSDK-UGU4PITWNI";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iscsiServers.listMetricDefinition(
    deviceName,
    iscsiServerName,
    resourceGroupName,
    managerName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

iscsiServersListMetricDefinition().catch(console.error);
