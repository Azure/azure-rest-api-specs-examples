
/**
 * Samples for FederatedIdentityCredentials List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2024-11-30/examples/
     * FederatedIdentityCredentialList.json
     */
    /**
     * Sample code: FederatedIdentityCredentialList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void federatedIdentityCredentialList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.identities().manager().serviceClient().getFederatedIdentityCredentials().list("rgName", "resourceName",
            null, null, com.azure.core.util.Context.NONE);
    }
}
