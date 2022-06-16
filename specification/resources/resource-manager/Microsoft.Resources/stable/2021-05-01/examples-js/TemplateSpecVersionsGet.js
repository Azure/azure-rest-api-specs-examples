const { TemplateSpecsClient } = require("@azure/arm-templatespecs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Template Spec version from a specific Template Spec.
 *
 * @summary Gets a Template Spec version from a specific Template Spec.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecVersionsGet.json
 */
async function templateSpecVersionsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "templateSpecRG";
  const templateSpecName = "simpleTemplateSpec";
  const templateSpecVersion = "v1.0";
  const credential = new DefaultAzureCredential();
  const client = new TemplateSpecsClient(credential, subscriptionId);
  const result = await client.templateSpecVersions.get(
    resourceGroupName,
    templateSpecName,
    templateSpecVersion
  );
  console.log(result);
}

templateSpecVersionsGet().catch(console.error);
