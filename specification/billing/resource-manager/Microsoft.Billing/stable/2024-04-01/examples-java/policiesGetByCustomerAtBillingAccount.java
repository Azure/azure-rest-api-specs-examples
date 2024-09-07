
/**
 * Samples for Policies GetByCustomerAtBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * policiesGetByCustomerAtBillingAccount.json
     */
    /**
     * Sample code: PoliciesGetByCustomerAtBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void policiesGetByCustomerAtBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.policies().getByCustomerAtBillingAccountWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "11111111-1111-1111-1111-111111111111", com.azure.core.util.Context.NONE);
    }
}
