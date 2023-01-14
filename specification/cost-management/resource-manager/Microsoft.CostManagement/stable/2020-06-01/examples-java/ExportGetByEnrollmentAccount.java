/** Samples for Exports Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportGetByEnrollmentAccount.json
     */
    /**
     * Sample code: ExportGetByEnrollmentAccount.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportGetByEnrollmentAccount(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .getWithResponse(
                "providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456",
                "TestExport",
                null,
                com.azure.core.util.Context.NONE);
    }
}
