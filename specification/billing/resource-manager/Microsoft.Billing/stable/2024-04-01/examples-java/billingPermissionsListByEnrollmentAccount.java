
/**
 * Samples for BillingPermissions ListByEnrollmentAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingPermissionsListByEnrollmentAccount.json
     */
    /**
     * Sample code: BillingPermissionsListByEnrollmentAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingPermissionsListByEnrollmentAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingPermissions().listByEnrollmentAccount("6100092", "123456", com.azure.core.util.Context.NONE);
    }
}
