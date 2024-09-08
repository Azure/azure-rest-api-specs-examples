
/**
 * Samples for Departments Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/departmentGet.json
     */
    /**
     * Sample code: DepartmentGet.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void departmentGet(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.departments().getWithResponse("456598", "164821", com.azure.core.util.Context.NONE);
    }
}
