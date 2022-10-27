const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { getLongRunningPoller } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Create or update a cloud service. Please note some properties can be set only during cloud service creation.
 *
 * @summary Create or update a cloud service. Please note some properties can be set only during cloud service creation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-04-04/examples/CloudService_Create_WithSingleRoleAndRDP.json
 */
async function createNewCloudServiceWithSingleRoleAndRdpExtension() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const options = {
    body: {
      location: "westus",
      properties: {
        configuration: "{ServiceConfiguration}",
        extensionProfile: {
          extensions: [
            {
              name: "RDPExtension",
              properties: {
                type: "RDP",
                autoUpgradeMinorVersion: false,
                protectedSettings: "<PrivateConfig><Password>{password}</Password></PrivateConfig>",
                publisher: "Microsoft.Windows.Azure.Extensions",
                settings:
                  "<PublicConfig><UserName>UserAzure</UserName><Expiration>10/22/2021 15:05:45</Expiration></PublicConfig>",
                typeHandlerVersion: "1.2",
              },
            },
          ],
        },
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
          ],
        },
        upgradeMode: "Auto",
      },
    },
    queryParameters: { "api-version": "2022-04-04" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}",
      subscriptionId,
      resourceGroupName,
      cloudServiceName
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createNewCloudServiceWithSingleRoleAndRdpExtension().catch(console.error);
