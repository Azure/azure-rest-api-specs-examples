const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a export. Update operation requires latest eTag to be set in the request. You may obtain the latest eTag by performing a get operation. Create operation does not require eTag.
 *
 * @summary The operation to create or update a export. Update operation requires latest eTag to be set in the request. You may obtain the latest eTag by performing a get operation. Create operation does not require eTag.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportCreateOrUpdateByResourceGroup.json
 */
async function exportCreateOrUpdateByResourceGroup() {
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG";
  const exportName = "TestExport";
  const parameters = {
    format: "Csv",
    definition: {
      type: "ActualCost",
      dataSet: {
        configuration: {
          columns: ["Date", "MeterId", "ResourceId", "ResourceLocation", "Quantity"],
        },
        granularity: "Daily",
      },
      timeframe: "MonthToDate",
    },
    deliveryInfo: {
      destination: {
        container: "exports",
        resourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182",
        rootFolderPath: "ad-hoc",
      },
    },
    schedule: {
      recurrence: "Weekly",
      recurrencePeriod: {
        from: new Date("2020-06-01T00:00:00Z"),
        to: new Date("2020-10-31T00:00:00Z"),
      },
      status: "Active",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.createOrUpdate(scope, exportName, parameters);
  console.log(result);
}
