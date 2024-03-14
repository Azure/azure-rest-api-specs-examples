
import com.azure.resourcemanager.security.fluent.models.PricingInner;
import com.azure.resourcemanager.security.models.Enforce;
import com.azure.resourcemanager.security.models.PricingTier;

/**
 * Samples for Pricings Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/
     * PutPricingVMsByName_example.json
     */
    /**
     * Sample code: Update pricing on subscription (example for VirtualMachines plan).
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void updatePricingOnSubscriptionExampleForVirtualMachinesPlan(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.pricings().updateWithResponse("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", "VirtualMachines",
            new PricingInner().withPricingTier(PricingTier.STANDARD).withSubPlan("P2").withEnforce(Enforce.TRUE),
            com.azure.core.util.Context.NONE);
    }
}
