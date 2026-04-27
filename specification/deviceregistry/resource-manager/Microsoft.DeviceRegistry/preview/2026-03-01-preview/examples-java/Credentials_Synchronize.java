
/**
 * Samples for Credentials Synchronize.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Credentials_Synchronize.json
     */
    /**
     * Sample code: Credentials_Synchronize.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void credentialsSynchronize(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.credentials().synchronize("rgdeviceregistry", "mynamespace", com.azure.core.util.Context.NONE);
    }
}
