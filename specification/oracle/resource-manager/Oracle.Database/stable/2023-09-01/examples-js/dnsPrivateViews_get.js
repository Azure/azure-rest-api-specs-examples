const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a DnsPrivateView
 *
 * @summary Get a DnsPrivateView
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dnsPrivateViews_get.json
 */
async function getADnsPrivateViewByName() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const dnsprivateviewocid = "ocid1....aaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.dnsPrivateViews.get(location, dnsprivateviewocid);
  console.log(result);
}
