const { TemplateSpecsClient } = require("@azure/arm-templatespecs");
const { DefaultAzureCredential } = require("@azure/identity");

async function templateSpecsPatch() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "templateSpecRG";
  const templateSpecName = "simpleTemplateSpec";
  const templateSpecVersion = "v1.0";
  const templateSpecVersionUpdateModel = {
    tags: { myTag: "My Value" },
  };
  const options = {
    templateSpecVersionUpdateModel,
  };
  const credential = new DefaultAzureCredential();
  const client = new TemplateSpecsClient(credential, subscriptionId);
  const result = await client.templateSpecVersions.update(
    resourceGroupName,
    templateSpecName,
    templateSpecVersion,
    options,
  );
  console.log(result);
}

templateSpecsPatch().catch(console.error);
