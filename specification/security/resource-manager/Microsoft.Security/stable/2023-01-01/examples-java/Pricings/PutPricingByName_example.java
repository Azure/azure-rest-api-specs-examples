import com.azure.resourcemanager.security.fluent.models.PricingInner;
import com.azure.resourcemanager.security.models.PricingTier;

/** Samples for Pricings Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2023-01-01/examples/Pricings/PutPricingByName_example.json
     */
    /**
     * Sample code: Update pricing on subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void updatePricingOnSubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .pricings()
            .updateWithResponse(
                "CloudPosture",
                new PricingInner().withPricingTier(PricingTier.STANDARD),
                com.azure.core.util.Context.NONE);
    }
}
