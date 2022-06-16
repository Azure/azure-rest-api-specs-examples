const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the record sets of a specified type in a DNS zone.
 *
 * @summary Lists the record sets of a specified type in a DNS zone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListTXTRecordset.json
 */
async function listTxtRecordsets() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const recordType = "TXT";
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recordSets.listByType(resourceGroupName, zoneName, recordType)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listTxtRecordsets().catch(console.error);
