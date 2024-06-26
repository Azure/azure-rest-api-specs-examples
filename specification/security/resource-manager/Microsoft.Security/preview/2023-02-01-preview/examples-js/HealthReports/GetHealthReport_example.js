const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get health report of resource
 *
 * @summary Get health report of resource
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-02-01-preview/examples/HealthReports/GetHealthReport_example.json
 */
async function getHealthReportOfResource() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceId =
    "subscriptions/a1efb6ca-fbc5-4782-9aaa-5c7daded1ce2/resourcegroups/E2E-IBB0WX/providers/Microsoft.Security/securityconnectors/AwsConnectorAllOfferings";
  const healthReportName = "909c629a-bf39-4521-8e4f-10b443a0bc02";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.healthReportOperations.get(resourceId, healthReportName);
  console.log(result);
}
