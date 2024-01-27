
/** Samples for FederatedIdentityCredentials Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/
     * FederatedIdentityCredentialDelete.json
     */
    /**
     * Sample code: FederatedIdentityCredentialDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void federatedIdentityCredentialDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.identities().manager().serviceClient().getFederatedIdentityCredentials().deleteWithResponse("rgName",
            "resourceName", "ficResourceName", com.azure.core.util.Context.NONE);
    }
}
