const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a NewRelicMonitorResource
 *
 * @summary Create a NewRelicMonitorResource
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/Monitors_CreateOrUpdate_MaximumSet_Gen.json
 */
async function monitorsCreateOrUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NEWRELICOBSERVABILITY_RESOURCE_GROUP"] || "rgNewRelic";
  const monitorName = "cdlymktqw";
  const resource = {
    location: "k",
    planData: {
      billingCycle: "YEARLY",
      effectiveDate: new Date("2022-12-05T14:11:37.786Z"),
      planDetails: "tbbiaga",
      usageType: "PAYG",
    },
    userInfo: {
      country: "hslqnwdanrconqyekwbnttaetv",
      emailAddress: "%6%@4-g.N1.3F-kI1.Ue-.lJso",
      firstName: "vdftzcggirefejajwahhwhyibutramdaotvnuf",
      lastName: "bcsztgqovdlmzfkjdrngidwzqsevagexzzilnlc",
      phoneNumber: "krf",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NewRelicObservability(credential, subscriptionId);
  const result = await client.monitors.beginCreateOrUpdateAndWait(
    resourceGroupName,
    monitorName,
    resource,
  );
  console.log(result);
}
