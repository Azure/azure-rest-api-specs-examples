/** Samples for Agents Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/Agents_Get.json
     */
    /**
     * Sample code: Agents_Get.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void agentsGet(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .agents()
            .getWithResponse(
                "examples-rg", "examples-storageMoverName", "examples-agentName", com.azure.core.util.Context.NONE);
    }
}
