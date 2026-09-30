const { PlatformValidationClient } = require("@azure/arm-platformvalidation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list cloud validations by resource group
 *
 * @summary list cloud validations by resource group
 * x-ms-original-file: 2026-08-01-preview/CloudValidations_ListByResourceGroup_MaximumSet_Gen.json
 */
async function cloudValidationsListByResourceGroupMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new PlatformValidationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.cloudValidations.listByResourceGroup("rgvalidate")) {
    resArray.push(item);
  }

  console.log(resArray);
}
