const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a reservation order alias.
 *
 * @summary Create a reservation order alias.
 * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/ReservationOrderAliasCreate.json
 */
async function reservationOrderAliasCreate() {
  const reservationOrderAliasName = "reservationOrderAlias123";
  const body = {
    appliedScopeProperties: {
      resourceGroupId: "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg",
    },
    appliedScopeType: "Single",
    billingPlan: "P1M",
    billingScopeId: "/subscriptions/10000000-0000-0000-0000-000000000000",
    displayName: "ReservationOrder_2022-06-02",
    location: "eastus",
    quantity: 5,
    renew: true,
    reservedResourceProperties: { instanceFlexibility: "On" },
    reservedResourceType: "VirtualMachines",
    sku: { name: "Standard_M64s_v2" },
    term: "P1Y",
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.reservationOrderAlias.beginCreateAndWait(
    reservationOrderAliasName,
    body
  );
  console.log(result);
}

reservationOrderAliasCreate().catch(console.error);
