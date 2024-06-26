const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a MonitorResource
 *
 * @summary Update a MonitorResource
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/Monitors_Update_MaximumSet_Gen.json
 */
async function monitorsUpdateMaximumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const resource = {
    dynatraceEnvironmentProperties: {
      accountInfo: {},
      environmentInfo: {},
      singleSignOnProperties: {},
    },
    marketplaceSubscriptionStatus: "Active",
    monitoringStatus: "Enabled",
    planData: {
      billingCycle: "Monthly",
      effectiveDate: new Date("2019-08-30T15:14:33+02:00"),
      planDetails: "dynatraceapitestplan",
      usageType: "Committed",
    },
    tags: { environment: "Dev" },
    userInfo: {
      country: "westus2",
      emailAddress: "alice@microsoft.com",
      firstName: "Alice",
      lastName: "Bobab",
      phoneNumber: "123456",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.update(resourceGroupName, monitorName, resource);
  console.log(result);
}

monitorsUpdateMaximumSetGen().catch(console.error);
