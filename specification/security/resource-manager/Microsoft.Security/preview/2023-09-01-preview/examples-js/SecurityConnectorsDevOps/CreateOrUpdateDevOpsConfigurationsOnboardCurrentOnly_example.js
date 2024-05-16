const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a DevOps Configuration.
 *
 * @summary Creates or updates a DevOps Configuration.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/CreateOrUpdateDevOpsConfigurationsOnboardCurrentOnly_example.json
 */
async function createOrUpdateDevOpsConfigurationsOnboardCurrentOnly() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "0806e1cd-cfda-4ff8-b99c-2b0af42cffd3";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "myRg";
  const securityConnectorName = "mySecurityConnectorName";
  const devOpsConfiguration = {
    properties: {
      authorization: { code: "00000000000000000000" },
      autoDiscovery: "Disabled",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.devOpsConfigurations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    securityConnectorName,
    devOpsConfiguration,
  );
  console.log(result);
}
