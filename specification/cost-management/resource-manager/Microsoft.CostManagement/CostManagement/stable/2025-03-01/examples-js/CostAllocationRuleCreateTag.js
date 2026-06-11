const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create/Update a rule to allocate cost between different resources within a billing account or enterprise enrollment.
 *
 * @summary create/Update a rule to allocate cost between different resources within a billing account or enterprise enrollment.
 * x-ms-original-file: 2025-03-01/CostAllocationRuleCreateTag.json
 */
async function costAllocationRulesCreateTag() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.costAllocationRules.createOrUpdate("100", "testRule", {
    properties: {
      description: "This is a testRule",
      status: "Active",
      details: {
        sourceResources: [{ name: "category", resourceType: "Tag", values: ["devops"] }],
        targetResources: [
          {
            name: "ResourceGroupName",
            policyType: "FixedProportion",
            resourceType: "Dimension",
            values: [
              { name: "destinationRG", percentage: 33.33 },
              { name: "destinationRG2", percentage: 33.33 },
              { name: "destinationRG3", percentage: 33.34 },
            ],
          },
        ],
      },
    },
  });
  console.log(result);
}
