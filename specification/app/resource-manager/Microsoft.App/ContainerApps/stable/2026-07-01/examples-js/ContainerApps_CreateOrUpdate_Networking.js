const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Container App.
 *
 * @summary create or update a Container App.
 * x-ms-original-file: 2026-07-01/ContainerApps_CreateOrUpdate_Networking.json
 */
async function createOrUpdateContainerAppWithPerAppOutboundVNetSubnet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerApps.createOrUpdate("rg", "testcontainerApp0", {
    location: "East US",
    environmentId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/expressenv",
    networking: {
      outboundVnetSubnetId:
        "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/appsubnet",
    },
    template: {
      containers: [
        {
          name: "testcontainerApp0",
          image: "repo/testcontainerApp0:v1",
          resources: { cpu: 0.2, memory: "100Mi" },
        },
      ],
    },
  });
  console.log(result);
}
