/** Samples for Agents Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/Agents_Delete.json
     */
    /**
     * Sample code: Agents_Delete.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void agentsDelete(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .agents()
            .delete("examples-rg", "examples-storageMoverName", "examples-agentName", com.azure.core.util.Context.NONE);
    }
}
