/** Samples for Alerts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/BillingAccountAlerts.json
     */
    /**
     * Sample code: BillingAccountAlerts.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void billingAccountAlerts(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .alerts()
            .listWithResponse(
                "providers/Microsoft.Billing/billingAccounts/12345:6789", com.azure.core.util.Context.NONE);
    }
}
