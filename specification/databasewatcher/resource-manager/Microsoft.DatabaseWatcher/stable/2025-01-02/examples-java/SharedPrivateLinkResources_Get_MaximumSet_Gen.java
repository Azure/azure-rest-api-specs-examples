
/**
 * Samples for SharedPrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/SharedPrivateLinkResources_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: SharedPrivateLinkResources_Get_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void sharedPrivateLinkResourcesGetMaximumSet(
        com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.sharedPrivateLinkResources().getWithResponse("apiTest-ddat4p", "databasemo3ej9ih", "monitoringh22eed",
            com.azure.core.util.Context.NONE);
    }
}
