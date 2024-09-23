
/**
 * Samples for FabricCapacities ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-11-01/FabricCapacities_ListByResourceGroup.json
     */
    /**
     * Sample code: List capacities by resource group.
     * 
     * @param manager Entry point to FabricManager.
     */
    public static void listCapacitiesByResourceGroup(com.azure.resourcemanager.fabric.FabricManager manager) {
        manager.fabricCapacities().listByResourceGroup("TestRG", com.azure.core.util.Context.NONE);
    }
}
