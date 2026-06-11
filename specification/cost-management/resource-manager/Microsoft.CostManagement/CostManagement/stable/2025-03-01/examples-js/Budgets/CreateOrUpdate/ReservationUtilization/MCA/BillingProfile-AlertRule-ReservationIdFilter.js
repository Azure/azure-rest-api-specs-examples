const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a budget. You can optionally provide an eTag if desired as a form of concurrency control. To obtain the latest eTag for a given budget, perform a get operation prior to your put operation.
 *
 * @summary the operation to create or update a budget. You can optionally provide an eTag if desired as a form of concurrency control. To obtain the latest eTag for a given budget, perform a get operation prior to your put operation.
 * x-ms-original-file: 2025-03-01/Budgets/CreateOrUpdate/ReservationUtilization/MCA/BillingProfile-AlertRule-ReservationIdFilter.json
 */
async function createOrUpdateReservationUtilizationBillingProfileMCAAlertRuleReservationIdFilter() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.budgets.createOrUpdate(
    "providers/Microsoft.Billing/billingAccounts/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee:ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj_2023-04-01/billingProfiles/KKKK-LLLL-MMM-NNN",
    "TestAlertRule",
    {
      eTag: '"1d34d016a593709"',
      category: "ReservationUtilization",
      filter: {
        dimensions: {
          name: "ReservationId",
          operator: "In",
          values: [
            "00000000-0000-0000-0000-000000000000",
            "00000000-0000-0000-0000-000000000001",
            "00000000-0000-0000-0000-000000000002",
          ],
        },
      },
      notifications: {
        Actual_LessThan_99_Percent: {
          contactEmails: ["johndoe@contoso.com", "janesmith@contoso.com"],
          enabled: true,
          frequency: "Daily",
          locale: "en-us",
          operator: "LessThan",
          threshold: 99,
        },
      },
      timeGrain: "Last30Days",
      timePeriod: {
        endDate: new Date("2025-04-01T00:00:00Z"),
        startDate: new Date("2023-04-01T00:00:00Z"),
      },
    },
  );
  console.log(result);
}
