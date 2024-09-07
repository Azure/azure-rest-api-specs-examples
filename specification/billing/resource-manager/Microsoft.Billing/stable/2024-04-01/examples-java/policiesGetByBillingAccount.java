
/**
 * Samples for Policies GetByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/policiesGetByBillingAccount.
     * json
     */
    /**
     * Sample code: PoliciesGetByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void policiesGetByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.policies().getByBillingAccountWithResponse("1234567", com.azure.core.util.Context.NONE);
    }
}
