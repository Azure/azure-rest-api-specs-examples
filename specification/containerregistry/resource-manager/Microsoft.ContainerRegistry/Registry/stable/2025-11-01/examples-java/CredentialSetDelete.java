
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
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void credentialSetDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getCredentialSets().delete("myResourceGroup",
            "myRegistry", "myCredentialSet", com.azure.core.util.Context.NONE);
    }
}
