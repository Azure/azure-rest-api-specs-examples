/** Samples for Exports Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportRunByEnrollmentAccount.json
     */
    /**
     * Sample code: ExportRunByEnrollmentAccount.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportRunByEnrollmentAccount(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .executeWithResponse(
                "providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456",
                "TestExport",
                com.azure.core.util.Context.NONE);
    }
}
