
/**
 * Samples for Credentials Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Delete_Credentials.json
     */
    /**
     * Sample code: Delete_Credentials.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void deleteCredentials(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.credentials().delete("rgdeviceregistry", "mynamespace", com.azure.core.util.Context.NONE);
    }
}
