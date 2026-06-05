const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a MonitorResource
 *
 * @summary update a MonitorResource
 * x-ms-original-file: 2024-04-24/Monitors_Update_MaximumSet_Gen.json
 */
async function monitorsUpdateMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.update("myResourceGroup", "myMonitor", {
    properties: {
      planData: {
        billingCycle: "Monthly",
        effectiveDate: new Date("2019-08-30T15:14:33+02:00"),
        planDetails: "dynatraceapitestplan",
        usageType: "Committed",
      },
    },
    tags: { Environment: "Dev" },
  });
  console.log(result);
}
