const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a security connector
 *
 * @summary Updates a security connector
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-03-01-preview/examples/SecurityConnectors/PatchSecurityConnector_example.json
 */
async function updateASecurityConnector() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "a5caac9c-5c04-49af-b3d0-e204f40345d5";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "exampleResourceGroup";
  const securityConnectorName = "exampleSecurityConnectorName";
  const securityConnector = {
    environmentData: { environmentType: "AwsAccount" },
    environmentName: "AWS",
    etag: "etag value (must be supplied for update)",
    hierarchyIdentifier: "exampleHierarchyId",
    location: "Central US",
    offerings: [
      {
        nativeCloudConnection: {
          cloudRoleArn: "arn:aws:iam::00000000:role/ASCMonitor",
        },
        offeringType: "CspmMonitorAws",
      },
    ],
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securityConnectors.update(
    resourceGroupName,
    securityConnectorName,
    securityConnector
  );
  console.log(result);
}
