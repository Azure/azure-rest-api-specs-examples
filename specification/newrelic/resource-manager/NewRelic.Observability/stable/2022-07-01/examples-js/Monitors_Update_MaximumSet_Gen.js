const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a NewRelicMonitorResource
 *
 * @summary Update a NewRelicMonitorResource
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_Update_MaximumSet_Gen.json
 */
async function monitorsUpdateMaximumSetGen() {
  const subscriptionId = process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "hfmjmpyqgezxkp";
  const resourceGroupName = process.env["NEWRELICOBSERVABILITY_RESOURCE_GROUP"] || "rgNewRelic";
  const monitorName = "cdlymktqw";
  const properties = {
    accountCreationSource: "LIFTR",
    identity: { type: "None", userAssignedIdentities: { key8903: {} } },
    newRelicAccountProperties: {
      accountInfo: {
        accountId: "xhqmg",
        ingestionKey: "wltnimmhqt",
        region: "ljcf",
      },
      organizationInfo: { organizationId: "k" },
      singleSignOnProperties: {
        enterpriseAppId: "kwiwfz",
        provisioningState: "Accepted",
        singleSignOnState: "Initial",
        singleSignOnUrl: "kvseueuljsxmfwpqctz",
      },
      userId: "vcscxlncofcuduadesd",
    },
    orgCreationSource: "LIFTR",
    planData: {
      billingCycle: "YEARLY",
      effectiveDate: new Date("2022-12-05T14:11:37.786Z"),
      planDetails: "tbbiaga",
      usageType: "PAYG",
    },
    tags: { key164: "jqakdrrmmyzytqu" },
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
  const result = await client.monitors.update(resourceGroupName, monitorName, properties);
  console.log(result);
}
