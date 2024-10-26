const { MicrosoftElastic } = require("@azure/arm-elastic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a OpenAI integration rule for a given monitor resource.
 *
 * @summary Create or update a OpenAI integration rule for a given monitor resource.
 * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/OpenAI_CreateOrUpdate.json
 */
async function openAiCreateOrUpdate() {
  const subscriptionId =
    process.env["ELASTIC_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ELASTIC_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "myMonitor";
  const integrationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftElastic(credential, subscriptionId);
  const result = await client.openAI.createOrUpdate(
    resourceGroupName,
    monitorName,
    integrationName,
  );
  console.log(result);
}
