
/**
 * Samples for FabricCapacities ListSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-11-01/FabricCapacities_ListSkus.json
     */
    /**
     * Sample code: List eligible SKUs for a new capacity.
     * 
     * @param manager Entry point to FabricManager.
     */
    public static void listEligibleSKUsForANewCapacity(com.azure.resourcemanager.fabric.FabricManager manager) {
        manager.fabricCapacities().listSkus(com.azure.core.util.Context.NONE);
    }
}
