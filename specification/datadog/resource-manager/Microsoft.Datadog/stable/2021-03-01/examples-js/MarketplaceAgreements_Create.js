const { MicrosoftDatadogClient } = require("@azure/arm-datadog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Datadog marketplace agreement in the subscription.
 *
 * @summary Create Datadog marketplace agreement in the subscription.
 * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/MarketplaceAgreements_Create.json
 */
async function marketplaceAgreementsCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftDatadogClient(credential, subscriptionId);
  const result = await client.marketplaceAgreements.createOrUpdate();
  console.log(result);
}

marketplaceAgreementsCreateOrUpdate().catch(console.error);
