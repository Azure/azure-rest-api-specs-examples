const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the forecast charges for external cloud provider type defined.
 *
 * @summary Lists the forecast charges for external cloud provider type defined.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExternalSubscriptionForecast.json
 */
async function externalSubscriptionForecast() {
  const externalCloudProviderType = "externalSubscriptions";
  const externalCloudProviderId = "100";
  const parameters = {
    type: "Usage",
    dataset: {
      aggregation: { totalCost: { name: "Cost", function: "Sum" } },
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
    timePeriod: {
      from: new Date("2022-08-01T00:00:00+00:00"),
      to: new Date("2022-08-31T23:59:59+00:00"),
    },
    timeframe: "Custom",
  };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.forecast.externalCloudProviderUsage(
    externalCloudProviderType,
    externalCloudProviderId,
    parameters
  );
  console.log(result);
}
