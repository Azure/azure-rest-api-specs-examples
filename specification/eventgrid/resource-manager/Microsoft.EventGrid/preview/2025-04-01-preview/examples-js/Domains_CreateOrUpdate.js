const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Asynchronously creates or updates a new domain with the specified parameters.
 *
 * @summary Asynchronously creates or updates a new domain with the specified parameters.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Domains_CreateOrUpdate.json
 */
async function domainsCreateOrUpdate() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const domainName = "exampledomain1";
  const domainInfo = {
    inboundIpRules: [
      { action: "Allow", ipMask: "12.18.30.15" },
      { action: "Allow", ipMask: "12.18.176.1" },
    ],
    location: "westus2",
    publicNetworkAccess: "Enabled",
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.domains.beginCreateOrUpdateAndWait(
    resourceGroupName,
    domainName,
    domainInfo,
  );
  console.log(result);
}
