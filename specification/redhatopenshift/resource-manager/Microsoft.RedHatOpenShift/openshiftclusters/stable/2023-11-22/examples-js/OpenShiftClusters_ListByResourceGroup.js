const { AzureRedHatOpenShiftClient } = require("@azure/arm-redhatopenshift");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation returns properties of each OpenShift cluster.
 *
 * @summary The operation returns properties of each OpenShift cluster.
 * x-ms-original-file: specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/examples/OpenShiftClusters_ListByResourceGroup.json
 */
async function listsOpenShiftClustersInTheSpecifiedSubscriptionAndResourceGroup() {
  const subscriptionId = process.env["REDHATOPENSHIFT_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["REDHATOPENSHIFT_RESOURCE_GROUP"] || "resourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new AzureRedHatOpenShiftClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.openShiftClusters.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
