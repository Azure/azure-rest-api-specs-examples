/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/DepartmentDimensionsListExpandAndTop.json
     */
    /**
     * Sample code: DepartmentDimensionsListExpandAndTop-Legacy.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void departmentDimensionsListExpandAndTopLegacy(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "providers/Microsoft.Billing/billingAccounts/100/departments/123",
                null,
                "properties/data",
                null,
                5,
                com.azure.core.util.Context.NONE);
    }
}
