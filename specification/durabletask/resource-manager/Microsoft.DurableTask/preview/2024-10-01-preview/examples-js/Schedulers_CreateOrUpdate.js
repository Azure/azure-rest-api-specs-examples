const { DurableTaskClient } = require("@azure/arm-durabletask");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Scheduler
 *
 * @summary create or update a Scheduler
 * x-ms-original-file: 2024-10-01-preview/Schedulers_CreateOrUpdate.json
 */
async function schedulersCreateOrUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "EE9BD735-67CE-4A90-89C4-439D3F6A4C93";
  const client = new DurableTaskClient(credential, subscriptionId);
  const result = await client.schedulers.createOrUpdate("rgopenapi", "testscheduler", {
    location: "northcentralus",
    properties: { ipAllowlist: ["10.0.0.0/8"], sku: { name: "Dedicated" } },
    tags: { key7131: "ryohwcoiccwsnewjigfmijz", key2138: "fjaeecgnvqd" },
  });
  console.log(result);
}
