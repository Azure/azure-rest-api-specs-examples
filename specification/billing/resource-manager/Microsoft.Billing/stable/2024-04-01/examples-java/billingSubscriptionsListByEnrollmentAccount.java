
/**
 * Samples for BillingSubscriptions ListByEnrollmentAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingSubscriptionsListByEnrollmentAccount.json
     */
    /**
     * Sample code: BillingSubscriptionsListByEnrollmentAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingSubscriptionsListByEnrollmentAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptions().listByEnrollmentAccount("6564892", "172988", null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
