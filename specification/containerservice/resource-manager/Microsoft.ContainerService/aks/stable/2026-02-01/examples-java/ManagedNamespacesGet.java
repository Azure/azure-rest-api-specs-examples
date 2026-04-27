
/**
 * Samples for ManagedNamespaces Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedNamespacesGet.json
     */
    /**
     * Sample code: Get Managed Namespace.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getManagedNamespace(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedNamespaces().getWithResponse("rg1", "clustername1", "namespace1",
            com.azure.core.util.Context.NONE);
    }
}
