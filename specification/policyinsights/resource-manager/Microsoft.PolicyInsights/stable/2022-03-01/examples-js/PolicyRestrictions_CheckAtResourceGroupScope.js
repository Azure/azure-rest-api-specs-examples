const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks what restrictions Azure Policy will place on a resource within a resource group. Use this when the resource group the resource will be created in is already known.
 *
 * @summary Checks what restrictions Azure Policy will place on a resource within a resource group. Use this when the resource group the resource will be created in is already known.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-03-01/examples/PolicyRestrictions_CheckAtResourceGroupScope.json
 */
async function checkPolicyRestrictionsAtResourceGroupScope() {
  const subscriptionId = "35ee058e-5fa0-414c-8145-3ebb8d09b6e2";
  const resourceGroupName = "vmRg";
  const parameters = {
    pendingFields: [
      { field: "name", values: ["myVMName"] },
      {
        field: "location",
        values: ["eastus", "westus", "westus2", "westeurope"],
      },
      { field: "tags" },
    ],
    resourceDetails: {
      apiVersion: "2019-12-01",
      resourceContent: {
        type: "Microsoft.Compute/virtualMachines",
        properties: { priority: "Spot" },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.policyRestrictions.checkAtResourceGroupScope(
    resourceGroupName,
    parameters
  );
  console.log(result);
}

checkPolicyRestrictionsAtResourceGroupScope().catch(console.error);
