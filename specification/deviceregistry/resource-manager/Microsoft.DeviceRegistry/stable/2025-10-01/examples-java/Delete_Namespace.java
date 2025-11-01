
/**
 * Samples for Namespaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Delete_Namespace.json
     */
    /**
     * Sample code: Delete_Namespace.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void deleteNamespace(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaces().delete("myResourceGroup", "adr-namespace-gbk0925-n01", com.azure.core.util.Context.NONE);
    }
}
