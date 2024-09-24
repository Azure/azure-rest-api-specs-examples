
/**
 * Samples for FabricCapacities Suspend.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-11-01/FabricCapacities_Suspend.json
     */
    /**
     * Sample code: Suspend capacity.
     * 
     * @param manager Entry point to FabricManager.
     */
    public static void suspendCapacity(com.azure.resourcemanager.fabric.FabricManager manager) {
        manager.fabricCapacities().suspend("TestRG", "azsdktest", com.azure.core.util.Context.NONE);
    }
}
