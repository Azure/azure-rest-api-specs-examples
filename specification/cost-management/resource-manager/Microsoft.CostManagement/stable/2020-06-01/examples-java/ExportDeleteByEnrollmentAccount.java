/** Samples for Exports Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportDeleteByEnrollmentAccount.json
     */
    /**
     * Sample code: ExportDeleteByEnrollmentAccount.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportDeleteByEnrollmentAccount(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .deleteByResourceGroupWithResponse(
                "providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456",
                "TestExport",
                com.azure.core.util.Context.NONE);
    }
}
