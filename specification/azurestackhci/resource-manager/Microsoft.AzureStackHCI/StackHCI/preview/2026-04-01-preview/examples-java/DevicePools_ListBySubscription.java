
/**
 * Samples for DevicePools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/DevicePools_ListBySubscription.json
     */
    /**
     * Sample code: List device pools in a given subscription.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listDevicePoolsInAGivenSubscription(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.devicePools().list(com.azure.core.util.Context.NONE);
    }
}
