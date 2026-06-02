
/**
 * Samples for FederatedIdentityCredentials Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/FederatedIdentityCredentialDelete.json
     */
    /**
     * Sample code: FederatedIdentityCredentialDelete.
     * 
     * @param manager Entry point to MsiManager.
     */
    public static void federatedIdentityCredentialDelete(com.azure.resourcemanager.msi.MsiManager manager) {
        manager.serviceClient().getFederatedIdentityCredentials().deleteWithResponse("rgName", "resourceName",
            "ficResourceName", com.azure.core.util.Context.NONE);
    }
}
