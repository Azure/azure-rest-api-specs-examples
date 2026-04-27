
import com.azure.resourcemanager.deviceregistry.models.DeviceCredentialsRevokeRequest;

/**
 * Samples for NamespaceDevices Revoke.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/NamespaceDevices_RevokeFailure.json
     */
    /**
     * Sample code: NamespaceDevices_RevokeFailure.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        namespaceDevicesRevokeFailure(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDevices().revoke("rgdeviceregistry", "mynamespace", "device1",
            new DeviceCredentialsRevokeRequest().withDisable(true), com.azure.core.util.Context.NONE);
    }
}
