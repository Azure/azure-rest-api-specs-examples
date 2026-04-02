const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Maps Account. A Maps Account holds the keys which allow access to the Maps REST APIs.
 *
 * @summary create or update a Maps Account. A Maps Account holds the keys which allow access to the Maps REST APIs.
 * x-ms-original-file: 2025-10-01-preview/CreateAccountGen2.json
 */
async function createGen2Account() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
  const client = new AzureMapsManagementClient(credential, subscriptionId);
  const result = await client.accounts.createOrUpdate("myResourceGroup", "myMapsAccount", {
    kind: "Gen2",
    location: "eastus",
    properties: {
      cors: {
        corsRules: [{ allowedOrigins: ["http://www.contoso.com", "http://www.fabrikam.com"] }],
      },
      disableLocalAuth: true,
      locations: [{ locationName: "northeurope" }],
    },
    sku: { name: "G2" },
    tags: { test: "true" },
  });
  console.log(result);
}
