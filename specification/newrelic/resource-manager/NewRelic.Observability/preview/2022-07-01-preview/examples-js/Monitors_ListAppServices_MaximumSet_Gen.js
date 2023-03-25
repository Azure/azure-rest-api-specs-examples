const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the app service resources currently being monitored by the NewRelic resource.
 *
 * @summary List the app service resources currently being monitored by the NewRelic resource.
 * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/preview/2022-07-01-preview/examples/Monitors_ListAppServices_MaximumSet_Gen.json
 */
async function monitorsListAppServicesMaximumSetGen() {
  const subscriptionId = process.env["NEWRELICOBSERVABILITY_SUBSCRIPTION_ID"] || "nqmcgifgaqlf";
  const resourceGroupName = process.env["NEWRELICOBSERVABILITY_RESOURCE_GROUP"] || "rgNewRelic";
  const monitorName = "fhcjxnxumkdlgpwanewtkdnyuz";
  const request = {
    azureResourceIds: ["pvzrksrmzowobuhxpwiotnpcvjbu"],
    userEmail: "ruxvg@xqkmdhrnoo.hlmbpm",
  };
  const credential = new DefaultAzureCredential();
  const client = new NewRelicObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitors.listAppServices(resourceGroupName, monitorName, request)) {
    resArray.push(item);
  }
  console.log(resArray);
}
