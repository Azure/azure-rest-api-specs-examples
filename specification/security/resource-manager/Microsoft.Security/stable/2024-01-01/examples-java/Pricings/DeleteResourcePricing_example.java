
/**
 * Samples for Pricings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/
     * DeleteResourcePricing_example.json
     */
    /**
     * Sample code: Delete a pricing on resource.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteAPricingOnResource(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.pricings().deleteByResourceGroupWithResponse(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/DEMO/providers/Microsoft.Compute/virtualMachines/VM-1",
            "VirtualMachines", com.azure.core.util.Context.NONE);
    }
}
