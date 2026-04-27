
/**
 * Samples for PlatformUpdates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/PlatformUpdates_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: PlatformUpdates_Get_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        platformUpdatesGetMaximumSet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.platformUpdates().getWithResponse("westus2", "10.2408.0.1", com.azure.core.util.Context.NONE);
    }
}
