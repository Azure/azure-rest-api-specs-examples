
/**
 * Samples for Snapshots Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-06-01/examples/
     * ConfigurationStoresGetSnapshot.json
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
