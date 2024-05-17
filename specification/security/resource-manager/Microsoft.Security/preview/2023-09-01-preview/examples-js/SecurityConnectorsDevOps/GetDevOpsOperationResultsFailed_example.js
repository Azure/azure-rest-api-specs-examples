const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get devops long running operation result.
 *
 * @summary Get devops long running operation result.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/GetDevOpsOperationResultsFailed_example.json
 */
async function getDevOpsOperationResultsFailed() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "0806e1cd-cfda-4ff8-b99c-2b0af42cffd3";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "myRg";
  const securityConnectorName = "mySecurityConnectorName";
  const operationResultId = "8d4caace-e7b3-4b3e-af99-73f76829ebcf";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.devOpsOperationResults.get(
    resourceGroupName,
    securityConnectorName,
    operationResultId,
  );
  console.log(result);
}
