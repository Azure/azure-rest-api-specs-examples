/** Samples for Agents List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-03-01/examples/Agents_List.json
     */
    /**
     * Sample code: Agents_List.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void agentsList(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.agents().list("examples-rg", "examples-storageMoverName", com.azure.core.util.Context.NONE);
    }
}
