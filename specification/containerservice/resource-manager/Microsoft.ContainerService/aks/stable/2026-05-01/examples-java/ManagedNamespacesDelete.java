
/**
 * Samples for ManagedNamespaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/ManagedNamespacesDelete.json
     */
    /**
     * Sample code: Delete Managed Namespace.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        deleteManagedNamespace(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedNamespaces().delete("rg1", "clustername1", "namespace1",
            com.azure.core.util.Context.NONE);
    }
}
