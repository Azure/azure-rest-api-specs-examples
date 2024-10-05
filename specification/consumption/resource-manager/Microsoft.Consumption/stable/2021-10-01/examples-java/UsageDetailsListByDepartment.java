
/**
 * Samples for UsageDetails List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * UsageDetailsListByDepartment.json
     */
    /**
     * Sample code: DepartmentUsageDetailsList-Legacy.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void
        departmentUsageDetailsListLegacy(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.usageDetails().list("providers/Microsoft.Billing/Departments/1234", null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
