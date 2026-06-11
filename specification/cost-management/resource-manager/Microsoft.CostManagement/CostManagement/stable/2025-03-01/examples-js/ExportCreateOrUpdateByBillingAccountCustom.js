const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a export. Update operation requires latest eTag to be set in the request. You may obtain the latest eTag by performing a get operation. Create operation does not require eTag.
 *
 * @summary the operation to create or update a export. Update operation requires latest eTag to be set in the request. You may obtain the latest eTag by performing a get operation. Create operation does not require eTag.
 * x-ms-original-file: 2025-03-01/ExportCreateOrUpdateByBillingAccountCustom.json
 */
async function exportCreateOrUpdateByBillingAccountCustom() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.createOrUpdate(
    "providers/Microsoft.Billing/billingAccounts/123456",
    "TestExport",
    {
      identity: { type: "SystemAssigned" },
      location: "centralus",
      format: "Csv",
      compressionMode: "gzip",
      dataOverwriteBehavior: "OverwritePreviousReport",
      definition: {
        type: "ActualCost",
        dataSet: { configuration: { dataVersion: "2023-05-01" }, granularity: "Daily" },
        timePeriod: {
          from: new Date("2025-04-03T00:00:00.000Z"),
          to: new Date("2025-04-03T00:00:00.000Z"),
        },
        timeframe: "Custom",
      },
      deliveryInfo: {
        destination: {
          type: "AzureBlob",
          container: "exports",
          resourceId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182",
          rootFolderPath: "ad-hoc",
        },
      },
      exportDescription: "This is a test export.",
      partitionData: true,
      schedule: { status: "Inactive" },
    },
  );
  console.log(result);
}
