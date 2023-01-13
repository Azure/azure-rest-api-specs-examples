/** Samples for Exports List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportsGetByDepartment.json
     */
    /**
     * Sample code: ExportsGetByDepartment.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportsGetByDepartment(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .listWithResponse(
                "providers/Microsoft.Billing/billingAccounts/12/departments/123",
                null,
                com.azure.core.util.Context.NONE);
    }
}
