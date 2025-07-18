const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Uploads or downloads the list of domains for a DNS Resolver Domain List from a storage link.
 *
 * @summary Uploads or downloads the list of domains for a DNS Resolver Domain List from a storage link.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsResolverDomainList_BulkUpload.json
 */
async function uploadDnsResolverDomainListDomains() {
  const subscriptionId =
    process.env["DNSRESOLVER_SUBSCRIPTION_ID"] || "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = process.env["DNSRESOLVER_RESOURCE_GROUP"] || "sampleResourceGroup";
  const dnsResolverDomainListName = "sampleDnsResolverDomainList";
  const parameters = {
    action: "Upload",
    storageUrl:
      "https://sampleStorageAccount.blob.core.windows.net/sample-container/sampleBlob.txt?sv=2022-11-02&sr=b&sig=39Up9jzHkxhUIhFEjEh9594DJxe7w6cIRCgOV6ICGS0%3A377&sp=rcw",
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.dnsResolverDomainLists.beginBulkAndWait(
    resourceGroupName,
    dnsResolverDomainListName,
    parameters,
  );
  console.log(result);
}
