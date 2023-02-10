/** Samples for Alerts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/EnrollmentAccountAlerts.json
     */
    /**
     * Sample code: EnrollmentAccountAlerts.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void enrollmentAccountAlerts(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .alerts()
            .listWithResponse(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/enrollmentAccounts/456",
                com.azure.core.util.Context.NONE);
    }
}
