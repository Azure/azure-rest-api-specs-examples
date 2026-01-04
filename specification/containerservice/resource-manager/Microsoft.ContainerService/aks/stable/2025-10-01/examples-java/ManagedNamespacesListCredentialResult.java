
/**
 * Samples for ManagedNamespaces ListCredential.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedNamespacesListCredentialResult.json
     */
    /**
     * Sample code: List managed namespace credentials.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedNamespaceCredentials(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedNamespaces().listCredentialWithResponse("rg1",
            "clustername1", "namespace1", com.azure.core.util.Context.NONE);
    }
}
