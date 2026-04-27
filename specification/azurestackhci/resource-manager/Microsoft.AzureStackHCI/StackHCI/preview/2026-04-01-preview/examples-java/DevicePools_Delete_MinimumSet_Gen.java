
/**
 * Samples for DevicePools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/DevicePools_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: DevicePools_Delete_MinimumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        devicePoolsDeleteMinimumSet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.devicePools().delete("ArcInstance-rg", "devicePool1", com.azure.core.util.Context.NONE);
    }
}
