const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all run outputs for the specified Image Template resource
 *
 * @summary List all run outputs for the specified Image Template resource
 * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2024-02-01/examples/ListRunOutputs.json
 */
async function retrieveAListOfAllOutputsCreatedByTheLastRunOfAnImageTemplate() {
  const subscriptionId = process.env["IMAGEBUILDER_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["IMAGEBUILDER_RESOURCE_GROUP"] || "myResourceGroup";
  const imageTemplateName = "myImageTemplate";
  const credential = new DefaultAzureCredential();
  const client = new ImageBuilderClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineImageTemplates.listRunOutputs(
    resourceGroupName,
    imageTemplateName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
