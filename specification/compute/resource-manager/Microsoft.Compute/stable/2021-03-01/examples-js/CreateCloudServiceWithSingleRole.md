Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a cloud service. Please note some properties can be set only during cloud service creation.
 *
 * @summary Create or update a cloud service. Please note some properties can be set only during cloud service creation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/CreateCloudServiceWithSingleRole.json
 */
async function createNewCloudServiceWithSingleRole() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const parameters = {
    location: "westus",
    properties: {
      configuration: "{ServiceConfiguration}",
      networkProfile: {
        loadBalancerConfigurations: [
          {
            name: "myLoadBalancer",
            properties: {
              frontendIPConfigurations: [
                {
                  name: "myfe",
                  properties: {
                    publicIPAddress: {
                      id: "/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Network/publicIPAddresses/myPublicIP",
                    },
                  },
                },
              ],
            },
          },
        ],
      },
      packageUrl: "{PackageUrl}",
      roleProfile: {
        roles: [
          {
            name: "ContosoFrontend",
            sku: { name: "Standard_D1_v2", capacity: 1, tier: "Standard" },
          },
        ],
      },
      upgradeMode: "Auto",
    },
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServices.beginCreateOrUpdateAndWait(
    resourceGroupName,
    cloudServiceName,
    options
  );
  console.log(result);
}

createNewCloudServiceWithSingleRole().catch(console.error);
```
