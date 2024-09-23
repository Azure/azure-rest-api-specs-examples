
/**
 * Samples for FabricCapacities Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-11-01/FabricCapacities_Delete.json
     */
    /**
     * Sample code: Delete a capacity.
     * 
     * @param manager Entry point to FabricManager.
     */
    public static void deleteACapacity(com.azure.resourcemanager.fabric.FabricManager manager) {
        manager.fabricCapacities().delete("TestRG", "azsdktest", com.azure.core.util.Context.NONE);
    }
}
