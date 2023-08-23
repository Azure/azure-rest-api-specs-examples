const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Maps Account. A Maps Account holds the keys which allow access to the Maps REST APIs.
 *
 * @summary Create or update a Maps Account. A Maps Account holds the keys which allow access to the Maps REST APIs.
 * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/CreateAccountGen2.json
 */
async function createGen2Account() {
  const subscriptionId =
    process.env["MAPS_SUBSCRIPTION_ID"] || "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
  const resourceGroupName = process.env["MAPS_RESOURCE_GROUP"] || "myResourceGroup";
  const accountName = "myMapsAccount";
  const mapsAccount = {
    kind: "Gen2",
    location: "eastus",
    properties: {
      cors: {
        corsRules: [
          {
            allowedOrigins: ["http://www.contoso.com", "http://www.fabrikam.com"],
          },
        ],
      },
      disableLocalAuth: true,
    },
    sku: { name: "G2" },
    tags: { test: "true" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMapsManagementClient(credential, subscriptionId);
  const result = await client.accounts.createOrUpdate(resourceGroupName, accountName, mapsAccount);
  console.log(result);
}
