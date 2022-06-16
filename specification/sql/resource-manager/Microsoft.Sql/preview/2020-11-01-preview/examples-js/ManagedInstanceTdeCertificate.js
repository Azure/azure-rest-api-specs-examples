const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a TDE certificate for a given server.
 *
 * @summary Creates a TDE certificate for a given server.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceTdeCertificate.json
 */
async function uploadATdeCertificate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000001";
  const resourceGroupName = "testtdecert";
  const managedInstanceName = "testtdecert";
  const parameters = { privateBlob: "MIIXXXXXXXX" };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstanceTdeCertificates.beginCreateAndWait(
    resourceGroupName,
    managedInstanceName,
    parameters
  );
  console.log(result);
}

uploadATdeCertificate().catch(console.error);
