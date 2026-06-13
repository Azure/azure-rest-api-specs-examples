const { ImageBuilderClient } = require("@azure/arm-imagebuilder");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update the tags for this Virtual Machine Image Template
 *
 * @summary update the tags for this Virtual Machine Image Template
 * x-ms-original-file: 2025-10-01/UpdateImageTemplateToRemoveIdentities.json
 */
async function removeIdentitiesForAnImageTemplate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ImageBuilderClient(credential, subscriptionId);
  const result = await client.virtualMachineImageTemplates.update(
    "myResourceGroup",
    "myImageTemplate",
    { identity: { type: "None" } },
  );
  console.log(result);
}
