const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a certificate used on the endpoint with the given id.
 *
 * @summary Gets a certificate used on the endpoint with the given id.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/EndpointCertificatesGet.json
 */
async function getsAnEndpointCertificate() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "38e0dc56-907f-45ba-a97c-74233baad471";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "testrg";
  const managedInstanceName = "testcl";
  const endpointType = "DATABASE_MIRRORING";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.endpointCertificates.get(
    resourceGroupName,
    managedInstanceName,
    endpointType,
  );
  console.log(result);
}
