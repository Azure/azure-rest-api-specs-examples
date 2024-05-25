const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a DnsPrivateZone
 *
 * @summary Get a DnsPrivateZone
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/dnsPrivateZones_get.json
 */
async function getADnsPrivateZoneByName() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const dnsprivatezonename = "example-dns-private-zone";
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.dnsPrivateZones.get(location, dnsprivatezonename);
  console.log(result);
}
