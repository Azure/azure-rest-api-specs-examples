const { TemplateSpecsClient } = require("@azure/arm-templatespecs");
const { DefaultAzureCredential } = require("@azure/identity");

async function templatesSpecsListBySubscription() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new TemplateSpecsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.templateSpecs.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

templatesSpecsListBySubscription().catch(console.error);
