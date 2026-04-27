
/**
 * Samples for Credentials ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/List_Credentials_ByResourceGroup.json
     */
    /**
     * Sample code: List_Credentials_ByResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listCredentialsByResourceGroup(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.credentials().listByResourceGroup("rgdeviceregistry", "mynamespace", com.azure.core.util.Context.NONE);
    }
}
