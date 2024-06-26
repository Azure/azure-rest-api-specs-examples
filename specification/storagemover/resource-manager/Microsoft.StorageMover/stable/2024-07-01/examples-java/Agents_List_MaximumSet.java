
/**
 * Samples for Agents List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/
     * Agents_List_MaximumSet.json
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
