/** Samples for Exports Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportRunByDepartment.json
     */
    /**
     * Sample code: ExportRunByDepartment.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportRunByDepartment(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .executeWithResponse(
                "providers/Microsoft.Billing/billingAccounts/12/departments/1234",
                "TestExport",
                com.azure.core.util.Context.NONE);
    }
}
