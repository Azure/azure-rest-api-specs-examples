const { AzureRedHatOpenShiftClient } = require("@azure/arm-redhatopenshift");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation returns the installable OpenShift versions as strings.
 *
 * @summary The operation returns the installable OpenShift versions as strings.
 * x-ms-original-file: specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/examples/OpenShiftVersions_List.json
 */
async function listsAllOpenShiftVersionsAvailableToInstallInTheSpecifiedLocation() {
  const subscriptionId = process.env["REDHATOPENSHIFT_SUBSCRIPTION_ID"] || "subscriptionId";
  const location = "location";
  const credential = new DefaultAzureCredential();
  const client = new AzureRedHatOpenShiftClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.openShiftVersions.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
