const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a Managed Environment used to host container apps.
 *
 * @summary creates or updates a Managed Environment used to host container apps.
 * x-ms-original-file: 2025-10-02-preview/ManagedEnvironments_CustomInfrastructureResourceGroup_Create.json
 */
async function createEnvironmentWithCustomInfrastructureResourceGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironments.createOrUpdate("examplerg", "testcontainerenv", {
    location: "East US",
    appLogsConfiguration: {
      logAnalyticsConfiguration: { customerId: "string", sharedKey: "string" },
    },
    availabilityZones: ["1", "2", "3"],
    customDomainConfiguration: {
      certificatePassword: "1234",
      certificateValue: Buffer.from("Y2VydA==", "base64"),
      dnsSuffix: "www.my-name.com",
    },
    daprAIConnectionString:
      "InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/",
    infrastructureResourceGroup: "myInfrastructureRgName",
    vnetConfiguration: {
      infrastructureSubnetId:
        "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/RGName/providers/Microsoft.Network/virtualNetworks/VNetName/subnets/subnetName1",
    },
    workloadProfiles: [
      {
        name: "My-GP-01",
        enableFips: true,
        maximumCount: 12,
        minimumCount: 3,
        workloadProfileType: "GeneralPurpose",
      },
      {
        name: "My-MO-01",
        maximumCount: 6,
        minimumCount: 3,
        workloadProfileType: "MemoryOptimized",
      },
      {
        name: "My-CO-01",
        maximumCount: 6,
        minimumCount: 3,
        workloadProfileType: "ComputeOptimized",
      },
      { name: "My-consumption-01", workloadProfileType: "Consumption" },
    ],
    zoneRedundant: true,
  });
  console.log(result);
}
