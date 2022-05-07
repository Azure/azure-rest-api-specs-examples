Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a security connector
 *
 * @summary Updates a security connector
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/SecurityConnectors/PatchSecurityConnector_example.json
 */
async function updateASecurityConnector() {
  const subscriptionId = "a5caac9c-5c04-49af-b3d0-e204f40345d5";
  const resourceGroupName = "exampleResourceGroup";
  const securityConnectorName = "exampleSecurityConnectorName";
  const securityConnector = {
    cloudName: "AWS",
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

updateASecurityConnector().catch(console.error);
```
