
/**
 * Samples for PlatformUpdates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/PlatformUpdates_ListByLocation_MaximumSet_Gen.json
     */
    /**
     * Sample code: PlatformUpdates_ListByLocation_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        platformUpdatesListByLocationMaximumSet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.platformUpdates().list("westus2", com.azure.core.util.Context.NONE);
    }
}
