import com.azure.core.util.Context;

/** Samples for Products ListByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ProductsListByBillingProfile.json
     */
    /**
     * Sample code: ProductsListByBillingProfile.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void productsListByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.products().listByBillingProfile("{billingAccountName}", "{billingProfileName}", null, Context.NONE);
    }
}
