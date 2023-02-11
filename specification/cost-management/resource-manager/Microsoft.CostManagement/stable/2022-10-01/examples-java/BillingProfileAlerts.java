/** Samples for Alerts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/BillingProfileAlerts.json
     */
    /**
     * Sample code: BillingProfileAlerts.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void billingProfileAlerts(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .alerts()
            .listWithResponse(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579",
                com.azure.core.util.Context.NONE);
    }
}
