const { MicrosoftElastic } = require("@azure/arm-elastic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fetch User API Key from internal database, if it was generated and stored while creating the Elasticsearch Organization.
 *
 * @summary Fetch User API Key from internal database, if it was generated and stored while creating the Elasticsearch Organization.
 * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/Organizations_GetApiKey.json
 */
async function organizationsGetApiKey() {
  const subscriptionId =
    process.env["ELASTIC_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftElastic(credential, subscriptionId);
  const result = await client.organizations.getApiKey();
  console.log(result);
}
