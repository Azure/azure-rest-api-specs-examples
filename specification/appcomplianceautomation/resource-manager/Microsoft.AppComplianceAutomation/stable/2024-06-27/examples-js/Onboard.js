const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Onboard given subscriptions to Microsoft.AppComplianceAutomation provider.
 *
 * @summary Onboard given subscriptions to Microsoft.AppComplianceAutomation provider.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Onboard.json
 */
async function onboard() {
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.providerActions.beginOnboardAndWait({
    subscriptionIds: [
      "00000000-0000-0000-0000-000000000000",
      "00000000-0000-0000-0000-000000000001",
    ],
  });
  console.log(result);
}
