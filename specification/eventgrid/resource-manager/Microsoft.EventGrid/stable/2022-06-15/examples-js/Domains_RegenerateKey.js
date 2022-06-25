const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate a shared access key for a domain.
 *
 * @summary Regenerate a shared access key for a domain.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/Domains_RegenerateKey.json
 */
async function domainsRegenerateKey() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const domainName = "exampledomain2";
  const regenerateKeyRequest = { keyName: "key1" };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.domains.regenerateKey(
    resourceGroupName,
    domainName,
    regenerateKeyRequest
  );
  console.log(result);
}

domainsRegenerateKey().catch(console.error);
