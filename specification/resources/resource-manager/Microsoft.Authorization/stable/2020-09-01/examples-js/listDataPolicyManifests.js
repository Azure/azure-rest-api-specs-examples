const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation retrieves a list of all the data policy manifests that match the optional given $filter. Valid values for $filter are: "$filter=namespace eq '{0}'". If $filter is not provided, the unfiltered list includes all data policy manifests for data resource types. If $filter=namespace is provided, the returned list only includes all data policy manifests that have a namespace matching the provided value.
 *
 * @summary This operation retrieves a list of all the data policy manifests that match the optional given $filter. Valid values for $filter are: "$filter=namespace eq '{0}'". If $filter is not provided, the unfiltered list includes all data policy manifests for data resource types. If $filter=namespace is provided, the returned list only includes all data policy manifests that have a namespace matching the provided value.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-09-01/examples/listDataPolicyManifests.json
 */
async function listDataPolicyManifests() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dataPolicyManifests.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDataPolicyManifests().catch(console.error);
