import com.azure.resourcemanager.billing.fluent.models.ProductInner;
import com.azure.resourcemanager.billing.models.AutoRenew;

/** Samples for Products Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateProduct.json
     */
    /**
     * Sample code: UpdateBillingProperty.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void updateBillingProperty(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .products()
            .updateWithResponse(
                "{billingAccountName}",
                "{productName}",
                new ProductInner().withAutoRenew(AutoRenew.OFF),
                com.azure.core.util.Context.NONE);
    }
}
