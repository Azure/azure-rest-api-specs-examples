const { TemplateSpecsClient } = require("@azure/arm-templatespecs");
const { DefaultAzureCredential } = require("@azure/identity");

async function templateSpecVersionsDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "templateSpecRG";
  const templateSpecName = "simpleTemplateSpec";
  const templateSpecVersion = "v1.0";
  const credential = new DefaultAzureCredential();
  const client = new TemplateSpecsClient(credential, subscriptionId);
  const result = await client.templateSpecVersions.delete(
    resourceGroupName,
    templateSpecName,
    templateSpecVersion,
  );
  console.log(result);
}

templateSpecVersionsDelete().catch(console.error);
