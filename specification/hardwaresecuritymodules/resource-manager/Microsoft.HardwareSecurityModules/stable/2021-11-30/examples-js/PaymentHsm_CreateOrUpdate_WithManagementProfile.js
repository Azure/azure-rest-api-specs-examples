const { AzureDedicatedHSMResourceProvider } = require("@azure/arm-hardwaresecuritymodules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or Update a dedicated HSM in the specified subscription.
 *
 * @summary Create or Update a dedicated HSM in the specified subscription.
 * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/PaymentHsm_CreateOrUpdate_WithManagementProfile.json
 */
async function createANewOrUpdateAnExistingPaymentHsmWithManagementProfile() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "hsm-group";
  const name = "hsm1";
  const parameters = {
    location: "westus",
    managementNetworkProfile: {
      networkInterfaces: [{ privateIpAddress: "1.0.0.2" }],
      subnet: {
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01",
      },
    },
    networkProfile: {
      networkInterfaces: [{ privateIpAddress: "1.0.0.1" }],
      subnet: {
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01",
      },
    },
    sku: { name: "payShield10K_LMK1_CPS60" },
    stampId: "stamp01",
    tags: { dept: "hsm", environment: "dogfood" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureDedicatedHSMResourceProvider(credential, subscriptionId);
  const result = await client.dedicatedHsmOperations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    name,
    parameters
  );
  console.log(result);
}

createANewOrUpdateAnExistingPaymentHsmWithManagementProfile().catch(console.error);
