import com.azure.core.util.Context;

/** Samples for Products Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/Product.json
     */
    /**
     * Sample code: Product.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void product(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.products().getWithResponse("{billingAccountName}", "{productName}", Context.NONE);
    }
}
