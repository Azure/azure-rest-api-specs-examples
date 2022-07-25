const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
 *
 * @summary Deletes a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/ProviderInstances_Delete.json
 */
async function deletesASapMonitorProvider() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "mySapMonitor";
  const providerInstanceName = "myProviderInstance";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.providerInstances.beginDeleteAndWait(
    resourceGroupName,
    monitorName,
    providerInstanceName
  );
  console.log(result);
}

deletesASapMonitorProvider().catch(console.error);
