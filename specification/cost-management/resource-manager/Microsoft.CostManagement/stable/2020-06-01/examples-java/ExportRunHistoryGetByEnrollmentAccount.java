/** Samples for Exports GetExecutionHistory. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportRunHistoryGetByEnrollmentAccount.json
     */
    /**
     * Sample code: ExportRunHistoryGetByEnrollmentAccount.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportRunHistoryGetByEnrollmentAccount(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .getExecutionHistoryWithResponse(
                "providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456",
                "TestExport",
                com.azure.core.util.Context.NONE);
    }
}
