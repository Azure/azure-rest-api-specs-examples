```java
import com.azure.core.util.Context;
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
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-billing_1.0.0-beta.2/sdk/billing/azure-resourcemanager-billing/README.md) on how to add the SDK to your project and authenticate.
