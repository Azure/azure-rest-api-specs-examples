import com.azure.core.util.Context;
import com.azure.resourcemanager.billing.fluent.models.BillingPropertyInner;

/** Samples for BillingProperty Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateBillingProperty.json
     */
    /**
     * Sample code: UpdateBillingProperty.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void updateBillingProperty(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingProperties().updateWithResponse(new BillingPropertyInner().withCostCenter("1010"), Context.NONE);
    }
}
