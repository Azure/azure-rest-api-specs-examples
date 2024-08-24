
/**
 * Samples for FederatedIdentityCredentials Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/
     * FederatedIdentityCredentialGet.json
     */
    /**
     * Sample code: FederatedIdentityCredentialGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void federatedIdentityCredentialGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.identities().manager().serviceClient().getFederatedIdentityCredentials().getWithResponse("rgName",
            "resourceName", "ficResourceName", com.azure.core.util.Context.NONE);
    }
}
