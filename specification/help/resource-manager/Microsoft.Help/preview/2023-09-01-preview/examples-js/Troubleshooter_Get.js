const { HelpRP } = require("@azure/arm-selfhelp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets troubleshooter instance result which includes the step status/result of the troubleshooter resource name that is being executed.<br/> Get API is used to retrieve the result of a Troubleshooter instance, which includes the status and result of each step in the Troubleshooter workflow. This API requires the Troubleshooter resource name that was created using the Create API.
 *
 * @summary Gets troubleshooter instance result which includes the step status/result of the troubleshooter resource name that is being executed.<br/> Get API is used to retrieve the result of a Troubleshooter instance, which includes the status and result of each step in the Troubleshooter workflow. This API requires the Troubleshooter resource name that was created using the Create API.
 * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/Troubleshooter_Get.json
 */
async function troubleshootersGet() {
  const scope =
    "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp";
  const troubleshooterName = "abf168ed-1b54-454a-86f6-e4b62253d3b1";
  const credential = new DefaultAzureCredential();
  const client = new HelpRP(credential);
  const result = await client.troubleshooters.get(scope, troubleshooterName);
  console.log(result);
}
