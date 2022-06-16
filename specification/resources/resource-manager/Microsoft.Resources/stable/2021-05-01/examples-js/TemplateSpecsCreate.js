const { TemplateSpecsClient } = require("@azure/arm-templatespecs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Template Spec.
 *
 * @summary Creates or updates a Template Spec.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecsCreate.json
 */
async function templateSpecsCreateUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "templateSpecRG";
  const templateSpecName = "simpleTemplateSpec";
  const templateSpec = {
    description: "A very simple Template Spec",
    location: "eastus",
  };
  const credential = new DefaultAzureCredential();
  const client = new TemplateSpecsClient(credential, subscriptionId);
  const result = await client.templateSpecs.createOrUpdate(
    resourceGroupName,
    templateSpecName,
    templateSpec
  );
  console.log(result);
}

templateSpecsCreateUpdate().catch(console.error);
