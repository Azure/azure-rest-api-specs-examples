
/**
 * Samples for CredentialSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/CredentialSetDelete.json
     */
    /**
     * Sample code: CredentialSetDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        credentialSetDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getCredentialSets().delete("myResourceGroup", "myRegistry", "myCredentialSet",
            com.azure.core.util.Context.NONE);
    }
}
