
/**
 * Samples for Namespaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Get_Namespace.json
     */
    /**
     * Sample code: Get_Namespace.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getNamespace(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaces().getByResourceGroupWithResponse("myResourceGroup", "adr-namespace-gbk0925-n01",
            com.azure.core.util.Context.NONE);
    }
}
