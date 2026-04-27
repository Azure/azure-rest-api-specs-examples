
/**
 * Samples for DevicePools ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/DevicePools_ListByResourceGroup.json
     */
    /**
     * Sample code: List device pools in a given resource group.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listDevicePoolsInAGivenResourceGroup(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.devicePools().listByResourceGroup("test-rg", com.azure.core.util.Context.NONE);
    }
}
