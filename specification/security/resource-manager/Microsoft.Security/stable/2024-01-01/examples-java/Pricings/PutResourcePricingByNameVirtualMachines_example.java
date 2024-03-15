
import com.azure.resourcemanager.security.fluent.models.PricingInner;
import com.azure.resourcemanager.security.models.PricingTier;

/**
 * Samples for Pricings Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/
     * PutResourcePricingByNameVirtualMachines_example.json
     */
    /**
     * Sample code: Update pricing on resource (example for VirtualMachines plan).
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void updatePricingOnResourceExampleForVirtualMachinesPlan(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.pricings().updateWithResponse(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/DEMO/providers/Microsoft.Compute/virtualMachines/VM-1",
            "virtualMachines", new PricingInner().withPricingTier(PricingTier.STANDARD).withSubPlan("P1"),
            com.azure.core.util.Context.NONE);
    }
}
