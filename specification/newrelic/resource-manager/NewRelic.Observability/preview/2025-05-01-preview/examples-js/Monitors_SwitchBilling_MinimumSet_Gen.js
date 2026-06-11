const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to switches the billing for the New Relic Monitor resource to be billed by Azure Marketplace
 *
 * @summary switches the billing for the New Relic Monitor resource to be billed by Azure Marketplace
 * x-ms-original-file: 2025-05-01-preview/Monitors_SwitchBilling_MinimumSet_Gen.json
 */
async function monitorsSwitchBillingMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NewRelicObservability(credential, subscriptionId);
  const result = await client.monitors.switchBilling("rgNewRelic", "fhcjxnxumkdlgpwanewtkdnyuz", {
    userEmail: "ruxvg@xqkmdhrnoo.hlmbpm",
  });
  console.log(result);
}
