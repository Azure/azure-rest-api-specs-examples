const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
 *
 * @summary Creates a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/ProviderInstances_Create_Root_Certificate.json
 */
async function createASapMonitorHanaProviderWithRootCertificate() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "mySapMonitor";
  const providerInstanceName = "myProviderInstance";
  const providerInstanceParameter = {
    providerSettings: {
      dbName: "db",
      dbPassword: "****",
      dbPasswordUri: "",
      dbUsername: "user",
      hostname: "name",
      instanceNumber: "00",
      providerType: "SapHana",
      sapSid: "SID",
      sqlPort: "0000",
      sslHostNameInCertificate: "xyz.domain.com",
      sslPreference: "RootCertificate",
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
