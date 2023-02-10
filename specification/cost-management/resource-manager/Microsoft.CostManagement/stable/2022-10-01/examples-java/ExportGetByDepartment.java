/** Samples for Exports Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportGetByDepartment.json
     */
    /**
     * Sample code: ExportGetByDepartment.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportGetByDepartment(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .getWithResponse(
                "providers/Microsoft.Billing/billingAccounts/12/departments/1234",
                "TestExport",
                null,
                com.azure.core.util.Context.NONE);
    }
}
