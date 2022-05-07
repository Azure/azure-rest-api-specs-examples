Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a security connector. If a security connector is already created and a subsequent request is issued for the same security connector id, then it will be updated.
 *
 * @summary Creates or updates a security connector. If a security connector is already created and a subsequent request is issued for the same security connector id, then it will be updated.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/SecurityConnectors/PutSecurityConnector_example.json
 */
async function createOrUpdateASecurityConnector() {
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
  const result = await client.securityConnectors.createOrUpdate(
    resourceGroupName,
    securityConnectorName,
    securityConnector
  );
  console.log(result);
}

createOrUpdateASecurityConnector().catch(console.error);
```
