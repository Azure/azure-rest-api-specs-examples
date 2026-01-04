
/**
 * Samples for ManagedNamespaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedNamespacesDelete.json
     */
    /**
     * Sample code: Delete Managed Namespace.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteManagedNamespace(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedNamespaces().delete("rg1", "clustername1",
            "namespace1", com.azure.core.util.Context.NONE);
    }
}
