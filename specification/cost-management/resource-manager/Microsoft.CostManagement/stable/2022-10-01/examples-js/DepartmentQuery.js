const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Query the usage data for scope defined.
 *
 * @summary Query the usage data for scope defined.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/DepartmentQuery.json
 */
async function departmentQueryLegacy() {
  const scope = "providers/Microsoft.Billing/billingAccounts/100/departments/123";
  const parameters = {
    type: "Usage",
    dataset: {
      filter: {
        and: [
          {
            or: [
              {
                dimensions: {
                  name: "ResourceLocation",
                  operator: "In",
                  values: ["East US", "West Europe"],
                },
              },
              {
                tags: {
                  name: "Environment",
                  operator: "In",
                  values: ["UAT", "Prod"],
                },
              },
            ],
          },
          {
            dimensions: {
              name: "ResourceGroup",
              operator: "In",
              values: ["API"],
            },
          },
        ],
      },
      granularity: "Daily",
    },
    timeframe: "MonthToDate",
  };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.query.usage(scope, parameters);
  console.log(result);
}
