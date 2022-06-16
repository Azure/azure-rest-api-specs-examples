import com.azure.core.util.Context;

/** Samples for Customers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/Customer.json
     */
    /**
     * Sample code: Customer.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void customer(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.customers().getWithResponse("{billingAccountName}", "{customerName}", null, Context.NONE);
    }
}
