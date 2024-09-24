
/**
 * Samples for FabricCapacities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-11-01/FabricCapacities_ListBySubscription.json
     */
    /**
     * Sample code: List capacities by subscription.
     * 
     * @param manager Entry point to FabricManager.
     */
    public static void listCapacitiesBySubscription(com.azure.resourcemanager.fabric.FabricManager manager) {
        manager.fabricCapacities().list(com.azure.core.util.Context.NONE);
    }
}
