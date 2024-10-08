
/**
 * Samples for BillingPermissions ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingPermissionsListByBillingAccount.json
     */
    /**
     * Sample code: BillingPermissionsListByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingPermissionsListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingPermissions().listByBillingAccount(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            com.azure.core.util.Context.NONE);
    }
}
