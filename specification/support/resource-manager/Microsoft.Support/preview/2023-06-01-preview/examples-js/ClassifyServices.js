const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Classify the list of right Azure services.
 *
 * @summary Classify the list of right Azure services.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/ClassifyServices.json
 */
async function classifyListOfAzureServices() {
  const serviceClassificationRequest = {
    issueSummary: "Can not connect to Windows VM",
    resourceId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgname/providers/Microsoft.Compute/virtualMachines/vmname",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const result = await client.serviceClassificationsNoSubscription.classifyServices(
    serviceClassificationRequest,
  );
  console.log(result);
}
