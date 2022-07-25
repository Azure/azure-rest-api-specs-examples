const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
 *
 * @summary Creates a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/Db2ProviderInstances_Create.json
 */
async function createADb2Provider() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "mySapMonitor";
  const providerInstanceName = "myProviderInstance";
  const providerInstanceParameter = {
    providerSettings: {
      dbName: "dbName",
      dbPassword: "password",
      dbPasswordUri: "",
      dbPort: "dbPort",
      dbUsername: "username",
      hostname: "hostname",
      providerType: "Db2",
      sapSid: "SID",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.providerInstances.beginCreateAndWait(
    resourceGroupName,
    monitorName,
    providerInstanceName,
    providerInstanceParameter
  );
  console.log(result);
}

createADb2Provider().catch(console.error);
