const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing LogProfilesResource. To update other fields use the CreateOrUpdate method.
 *
 * @summary Updates an existing LogProfilesResource. To update other fields use the CreateOrUpdate method.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/patchLogProfile.json
 */
async function patchALogProfile() {
  const subscriptionId = "df602c9c-7aa0-407d-a6fb-eb20c8bd1192";
  const logProfileName = "Rac46PostSwapRG";
  const logProfilesResource = {
    categories: ["Write", "Delete", "Action"],
    locations: ["global"],
    retentionPolicy: { days: 3, enabled: true },
    serviceBusRuleId: "",
    storageAccountId:
      "/subscriptions/df602c9c-7aa0-407d-a6fb-eb20c8bd1192/resourceGroups/JohnKemTest/providers/Microsoft.Storage/storageAccounts/johnkemtest8162",
    tags: { key1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.logProfiles.update(logProfileName, logProfilesResource);
  console.log(result);
}

patchALogProfile().catch(console.error);
