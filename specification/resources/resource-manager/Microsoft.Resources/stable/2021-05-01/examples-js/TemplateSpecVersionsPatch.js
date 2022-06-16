const { TemplateSpecsClient } = require("@azure/arm-templatespecs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates Template Spec Version tags with specified values.
 *
 * @summary Updates Template Spec Version tags with specified values.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecVersionsPatch.json
 */
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
    options
  );
  console.log(result);
}

templateSpecsPatch().catch(console.error);
