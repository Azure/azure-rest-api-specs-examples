
/**
 * Samples for Agents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/Agents_List_MaximumSet.json
     */
    /**
     * Sample code: Agents_List_MaximumSet.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void agentsListMaximumSet(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.agents().list("examples-rg", "examples-storageMoverName", com.azure.core.util.Context.NONE);
    }
}
