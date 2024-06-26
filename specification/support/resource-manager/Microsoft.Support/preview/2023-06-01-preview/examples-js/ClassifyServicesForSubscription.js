const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Classify the list of right Azure services.
 *
 * @summary Classify the list of right Azure services.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/ClassifyServicesForSubscription.json
 */
async function classifyListOfAzureServicesForASubscription() {
  const subscriptionId =
    process.env["SUPPORT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const serviceClassificationRequest = {
    issueSummary: "Can not connect to Windows VM",
    resourceId:
      "/subscriptions/76cb77fa-8b17-4eab-9493-b65dace99813/resourceGroups/rgname/providers/Microsoft.Compute/virtualMachines/vmname",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.serviceClassifications.classifyServices(serviceClassificationRequest);
  console.log(result);
}
