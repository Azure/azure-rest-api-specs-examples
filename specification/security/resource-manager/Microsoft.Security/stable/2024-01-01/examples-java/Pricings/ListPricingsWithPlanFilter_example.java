
/**
 * Samples for Pricings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/
     * ListPricingsWithPlanFilter_example.json
     */
    /**
     * Sample code: Get pricings on subscription with plans filter.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        getPricingsOnSubscriptionWithPlansFilter(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.pricings().listWithResponse("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
            "name in (VirtualMachines,KeyVaults)", com.azure.core.util.Context.NONE);
    }
}
