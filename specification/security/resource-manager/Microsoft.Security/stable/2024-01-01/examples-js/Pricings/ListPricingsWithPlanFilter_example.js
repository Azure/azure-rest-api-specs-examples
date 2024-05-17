const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists Microsoft Defender for Cloud pricing configurations of the scopeId, that match the optional given $filter. Valid scopes are: subscription id or a specific resource id (Supported resources are: 'VirtualMachines, VMSS and ARC Machines'). Valid $filter is: 'name in ({planName1},{planName2},...)'. If $filter is not provided, the unfiltered list will be returned. If '$filter=name in (planName1,planName2)' is provided, the returned list includes the pricings set for 'planName1' and 'planName2' only.
 *
 * @summary Lists Microsoft Defender for Cloud pricing configurations of the scopeId, that match the optional given $filter. Valid scopes are: subscription id or a specific resource id (Supported resources are: 'VirtualMachines, VMSS and ARC Machines'). Valid $filter is: 'name in ({planName1},{planName2},...)'. If $filter is not provided, the unfiltered list will be returned. If '$filter=name in (planName1,planName2)' is provided, the returned list includes the pricings set for 'planName1' and 'planName2' only.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/ListPricingsWithPlanFilter_example.json
 */
async function getPricingsOnSubscriptionWithPlansFilter() {
  const scopeId = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.pricings.list(scopeId);
  console.log(result);
}
