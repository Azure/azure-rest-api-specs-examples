const { NginxManagementClient } = require("@azure/arm-nginx");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update the NGINX deployment
 *
 * @summary Create or update the NGINX deployment
 * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/Deployments_Create.json
 */
async function deploymentsCreate() {
  const subscriptionId =
    process.env["NGINX_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NGINX_RESOURCE_GROUP"] || "myResourceGroup";
  const deploymentName = "myDeployment";
  const body = {
    name: "myDeployment",
    location: "West US",
    properties: {
      autoUpgradeProfile: { upgradeChannel: "stable" },
      managedResourceGroup: "myManagedResourceGroup",
      networkProfile: {
        frontEndIPConfiguration: {
          privateIPAddresses: [
            {
              privateIPAddress: "1.1.1.1",
              privateIPAllocationMethod: "Static",
              subnetId:
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet",
            },
          ],
          publicIPAddresses: [
            {
              id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/publicIPAddresses/myPublicIPAddress",
            },
          ],
        },
        networkInterfaceConfiguration: {
          subnetId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet",
        },
      },
      scalingProperties: { capacity: 10 },
      userProfile: { preferredEmail: "example@example.email" },
    },
    tags: { environment: "Dev" },
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new NginxManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    deploymentName,
    options,
  );
  console.log(result);
}
