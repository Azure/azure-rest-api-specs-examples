
/**
 * Samples for ManagedNamespaces ListCredential.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/ManagedNamespacesListCredentialResult.json
     */
    /**
     * Sample code: List managed namespace credentials.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listManagedNamespaceCredentials(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedNamespaces().listCredentialWithResponse("rg1", "clustername1", "namespace1",
            com.azure.core.util.Context.NONE);
    }
}
