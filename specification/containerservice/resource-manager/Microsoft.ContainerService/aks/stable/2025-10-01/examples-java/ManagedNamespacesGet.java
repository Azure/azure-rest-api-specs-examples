
/**
 * Samples for ManagedNamespaces Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedNamespacesGet.json
     */
    /**
     * Sample code: Get Managed Namespace.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedNamespace(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedNamespaces().getWithResponse("rg1",
            "clustername1", "namespace1", com.azure.core.util.Context.NONE);
    }
}
