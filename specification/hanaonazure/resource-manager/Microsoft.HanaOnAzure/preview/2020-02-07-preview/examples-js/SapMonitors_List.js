const { HanaManagementClient } = require("@azure/arm-hanaonazure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of SAP monitors in the specified subscription. The operations returns various properties of each SAP monitor.
 *
 * @summary Gets a list of SAP monitors in the specified subscription. The operations returns various properties of each SAP monitor.
 * x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_List.json
 */
async function listAllSapMonitorsInASubscription() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new HanaManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sapMonitors.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllSapMonitorsInASubscription().catch(console.error);
