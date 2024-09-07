
/**
 * Samples for BillingPermissions ListByDepartment.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingPermissionsListByDepartment.json
     */
    /**
     * Sample code: BillingPermissionsListByDepartment.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingPermissionsListByDepartment(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingPermissions().listByDepartment("6100092", "123456", com.azure.core.util.Context.NONE);
    }
}
