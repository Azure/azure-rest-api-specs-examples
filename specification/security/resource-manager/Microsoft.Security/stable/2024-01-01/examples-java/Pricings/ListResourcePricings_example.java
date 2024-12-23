
/**
 * Samples for Pricings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/
     * ListResourcePricings_example.json
     */
    /**
     * Sample code: Get pricings on resource.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getPricingsOnResource(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.pricings().listWithResponse(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/DEMO/providers/Microsoft.Compute/virtualMachines/VM-1",
            null, com.azure.core.util.Context.NONE);
    }
}
