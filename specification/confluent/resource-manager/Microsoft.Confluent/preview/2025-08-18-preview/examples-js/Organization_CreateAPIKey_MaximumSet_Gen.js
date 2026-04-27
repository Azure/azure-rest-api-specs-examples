const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates API key for a schema registry Cluster ID or Kafka Cluster ID under a environment
 *
 * @summary creates API key for a schema registry Cluster ID or Kafka Cluster ID under a environment
 * x-ms-original-file: 2025-08-18-preview/Organization_CreateAPIKey_MaximumSet_Gen.json
 */
async function organizationCreateAPIKeyMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "DC34558A-05D3-4370-AED8-75E60B381F94";
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.organization.createAPIKey(
    "rgconfluent",
    "pzvuoywx",
    "jqoxoahobqmhnklw",
    "ypyzlfhbml",
    { name: "izlvofweryqgdgq", description: "vdxsmrddjlsfcsnwjezjraxgbkn" },
  );
  console.log(result);
}
