
import com.azure.resourcemanager.billing.models.AutoRenew;
import com.azure.resourcemanager.billing.models.ProductPatch;
import com.azure.resourcemanager.billing.models.ProductProperties;

/**
 * Samples for Products Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/productsUpdate.json
     */
    /**
     * Sample code: ProductsUpdate.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void productsUpdate(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.products().updateWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "11111111-1111-1111-1111-111111111111",
            new ProductPatch().withProperties(new ProductProperties().withAutoRenew(AutoRenew.ON)),
            com.azure.core.util.Context.NONE);
    }
}
