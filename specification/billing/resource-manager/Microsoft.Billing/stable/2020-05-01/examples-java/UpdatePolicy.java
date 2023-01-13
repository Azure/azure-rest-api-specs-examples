import com.azure.resourcemanager.billing.fluent.models.PolicyInner;
import com.azure.resourcemanager.billing.models.MarketplacePurchasesPolicy;
import com.azure.resourcemanager.billing.models.ReservationPurchasesPolicy;
import com.azure.resourcemanager.billing.models.ViewChargesPolicy;

/** Samples for Policies Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdatePolicy.json
     */
    /**
     * Sample code: UpdatePolicy.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void updatePolicy(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .policies()
            .updateWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                new PolicyInner()
                    .withMarketplacePurchases(MarketplacePurchasesPolicy.ONLY_FREE_ALLOWED)
                    .withReservationPurchases(ReservationPurchasesPolicy.NOT_ALLOWED)
                    .withViewCharges(ViewChargesPolicy.ALLOWED),
                com.azure.core.util.Context.NONE);
    }
}
