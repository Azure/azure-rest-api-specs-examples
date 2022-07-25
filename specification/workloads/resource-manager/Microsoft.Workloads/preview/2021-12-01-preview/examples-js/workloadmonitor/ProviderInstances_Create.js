const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
 *
 * @summary Creates a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/ProviderInstances_Create.json
 */
async function createASapMonitorHanaProvider() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "mySapMonitor";
  const providerInstanceName = "myProviderInstance";
  const providerInstanceParameter = {
    providerSettings: {
      dbName: "db",
      dbPassword: "****",
      dbPasswordUri: "",
      dbSslCertificateUri: "https://storageaccount.blob.core.windows.net/containername/filename",
      dbUsername: "user",
      hostname: "name",
      instanceNumber: "00",
      providerType: "SapHana",
      sqlPort: "0000",
      sslHostNameInCertificate: "xyz.domain.com",
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

createASapMonitorHanaProvider().catch(console.error);
