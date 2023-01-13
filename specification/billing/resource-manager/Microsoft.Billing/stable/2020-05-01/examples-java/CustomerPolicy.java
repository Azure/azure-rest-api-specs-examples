/** Samples for Policies GetByCustomer. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomerPolicy.json
     */
    /**
     * Sample code: PolicyByCustomer.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void policyByCustomer(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .policies()
            .getByCustomerWithResponse("{billingAccountName}", "{customerName}", com.azure.core.util.Context.NONE);
    }
}
