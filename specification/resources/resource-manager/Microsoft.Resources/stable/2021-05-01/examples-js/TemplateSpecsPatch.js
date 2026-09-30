const { TemplateSpecsClient } = require("@azure/arm-templatespecs");
const { DefaultAzureCredential } = require("@azure/identity");

async function templateSpecsPatch() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "templateSpecRG";
  const templateSpecName = "simpleTemplateSpec";
  const templateSpec = { tags: { myTag: "My Value" } };
  const options = { templateSpec };
  const credential = new DefaultAzureCredential();
  const client = new TemplateSpecsClient(credential, subscriptionId);
  const result = await client.templateSpecs.update(resourceGroupName, templateSpecName, options);
  console.log(result);
}

templateSpecsPatch().catch(console.error);
