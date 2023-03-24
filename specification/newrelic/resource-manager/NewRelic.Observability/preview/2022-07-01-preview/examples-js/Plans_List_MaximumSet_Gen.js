const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List plans data
 *
 * @summary List plans data
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/preview/2022-07-01-preview/examples/Plans_List_MaximumSet_Gen.json
 */
async function plansListMaximumSetGen() {
  const subscriptionId = process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "hfmjmpyqgezxkp";
  const accountId = "pwuxgvrmkk";
  const organizationId = "hilawwjz";
  const options = { accountId, organizationId };
  const credential = new DefaultAzureCredential();
  const client = new NewRelicObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.plans.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
