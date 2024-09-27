
/**
 * Samples for StandbyContainerGroupPoolRuntimeViews Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/StandbyContainerGroupPoolRuntimeViews_Get.json
     */
    /**
     * Sample code: StandbyContainerGroupPoolRuntimeViews_Get.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void
        standbyContainerGroupPoolRuntimeViewsGet(com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyContainerGroupPoolRuntimeViews().getWithResponse("rgstandbypool", "pool", "latest",
            com.azure.core.util.Context.NONE);
    }
}
