
/**
 * Samples for Namespaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_Namespace_ByResourceGroup.json
     */
    /**
     * Sample code: List_Namespaces_ByResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listNamespacesByResourceGroup(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaces().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
