
/**
 * Samples for Namespaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_Namespace_BySubscription.json
     */
    /**
     * Sample code: List_Namespaces_BySubscription.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listNamespacesBySubscription(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaces().list(com.azure.core.util.Context.NONE);
    }
}
