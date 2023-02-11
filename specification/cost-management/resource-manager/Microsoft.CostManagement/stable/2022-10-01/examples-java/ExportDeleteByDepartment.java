/** Samples for Exports Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportDeleteByDepartment.json
     */
    /**
     * Sample code: ExportDeleteByDepartment.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportDeleteByDepartment(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .deleteByResourceGroupWithResponse(
                "providers/Microsoft.Billing/billingAccounts/12/departments/1234",
                "TestExport",
                com.azure.core.util.Context.NONE);
    }
}
