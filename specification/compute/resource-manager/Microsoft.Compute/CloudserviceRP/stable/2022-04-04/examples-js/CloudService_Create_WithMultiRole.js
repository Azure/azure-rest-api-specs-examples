const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a cloud service. Please note some properties can be set only during cloud service creation.
 *
 * @summary Create or update a cloud service. Please note some properties can be set only during cloud service creation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-04-04/examples/CloudService_Create_WithMultiRole.json
 */
async function createNewCloudServiceWithMultipleRoles() {
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
            name: "contosolb",
            properties: {
              frontendIPConfigurations: [
                {
                  name: "contosofe",
                  properties: {
                    publicIPAddress: {
                      id: "/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Network/publicIPAddresses/contosopublicip",
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
          {
            name: "ContosoBackend",
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

createNewCloudServiceWithMultipleRoles().catch(console.error);
