const { TemplateSpecsClient } = require("@azure/arm-templatespecs");
const { DefaultAzureCredential } = require("@azure/identity");

async function templateSpecVersionsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "templateSpecRG";
  const templateSpecName = "simpleTemplateSpec";
  const credential = new DefaultAzureCredential();
  const client = new TemplateSpecsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.templateSpecVersions.list(resourceGroupName, templateSpecName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

templateSpecVersionsList().catch(console.error);
