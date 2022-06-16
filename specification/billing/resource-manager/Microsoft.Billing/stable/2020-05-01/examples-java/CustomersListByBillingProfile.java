import com.azure.core.util.Context;

/** Samples for Customers ListByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomersListByBillingProfile.json
     */
    /**
     * Sample code: CustomersListByBillingProfile.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void customersListByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .customers()
            .listByBillingProfile("{billingAccountName}", "{billingProfileName}", null, null, Context.NONE);
    }
}
