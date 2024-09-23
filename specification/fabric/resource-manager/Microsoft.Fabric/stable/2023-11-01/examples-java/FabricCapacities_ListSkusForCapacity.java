
/**
 * Samples for FabricCapacities ListSkusForCapacity.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-11-01/FabricCapacities_ListSkusForCapacity.json
     */
    /**
     * Sample code: List eligible SKUs for an existing capacity.
     * 
     * @param manager Entry point to FabricManager.
     */
    public static void listEligibleSKUsForAnExistingCapacity(com.azure.resourcemanager.fabric.FabricManager manager) {
        manager.fabricCapacities().listSkusForCapacity("TestRG", "azsdktest", com.azure.core.util.Context.NONE);
    }
}
