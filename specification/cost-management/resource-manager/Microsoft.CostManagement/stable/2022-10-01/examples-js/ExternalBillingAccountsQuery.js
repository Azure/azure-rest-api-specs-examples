const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Query the usage data for external cloud provider type defined.
 *
 * @summary Query the usage data for external cloud provider type defined.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExternalBillingAccountsQuery.json
 */
async function externalBillingAccountQueryList() {
  const externalCloudProviderType = "externalBillingAccounts";
  const externalCloudProviderId = "100";
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
  const result = await client.query.usageByExternalCloudProviderType(
    externalCloudProviderType,
    externalCloudProviderId,
    parameters
  );
  console.log(result);
}
