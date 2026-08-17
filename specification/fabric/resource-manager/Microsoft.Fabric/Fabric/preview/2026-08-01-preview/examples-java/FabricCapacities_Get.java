
/**
 * Samples for FabricCapacities GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/FabricCapacities_Get.json
     */
    /**
     * Sample code: Get a capacity.
     * 
     * @param manager Entry point to FabricManager.
     */
    public static void getACapacity(com.azure.resourcemanager.fabric.FabricManager manager) {
        manager.fabricCapacities().getByResourceGroupWithResponse("TestRG", "azsdktest",
            com.azure.core.util.Context.NONE);
    }
}
