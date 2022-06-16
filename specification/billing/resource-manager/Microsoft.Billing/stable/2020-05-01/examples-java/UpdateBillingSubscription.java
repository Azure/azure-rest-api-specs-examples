import com.azure.core.util.Context;
import com.azure.resourcemanager.billing.fluent.models.BillingSubscriptionInner;

/** Samples for BillingSubscriptions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateBillingSubscription.json
     */
    /**
     * Sample code: UpdateBillingProperty.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void updateBillingProperty(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingSubscriptions()
            .updateWithResponse(
                "{billingAccountName}", new BillingSubscriptionInner().withCostCenter("ABC1234"), Context.NONE);
    }
}
