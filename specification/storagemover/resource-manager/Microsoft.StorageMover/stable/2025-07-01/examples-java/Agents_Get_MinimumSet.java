
/**
 * Samples for Agents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/Agents_Get_MinimumSet.json
     */
    /**
     * Sample code: Agents_Get_MinimumSet.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void agentsGetMinimumSet(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.agents().getWithResponse("examples-rg", "examples-storageMoverName", "examples-agentName",
            com.azure.core.util.Context.NONE);
    }
}
