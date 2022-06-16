import com.azure.core.util.Context;

/** Samples for Customers ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomersListByBillingAccount.json
     */
    /**
     * Sample code: CustomersListByBillingAccount.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void customersListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.customers().listByBillingAccount("{billingAccountName}", null, null, Context.NONE);
    }
}
