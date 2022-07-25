const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets properties of a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
 *
 * @summary Gets properties of a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/Db2ProviderInstances_Get.json
 */
async function getPropertiesOfADb2Provider() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "mySapMonitor";
  const providerInstanceName = "myProviderInstance";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.providerInstances.get(
    resourceGroupName,
    monitorName,
    providerInstanceName
  );
  console.log(result);
}

getPropertiesOfADb2Provider().catch(console.error);
