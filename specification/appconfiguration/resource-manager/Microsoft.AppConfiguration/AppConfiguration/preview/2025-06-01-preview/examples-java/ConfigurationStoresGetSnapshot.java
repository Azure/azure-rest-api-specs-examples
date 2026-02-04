
/**
 * Samples for Snapshots Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresGetSnapshot.json
     */
    /**
     * Sample code: Snapshots_Get.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void snapshotsGet(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.snapshots().getWithResponse("myResourceGroup", "contoso", "mySnapshot",
            com.azure.core.util.Context.NONE);
    }
}
