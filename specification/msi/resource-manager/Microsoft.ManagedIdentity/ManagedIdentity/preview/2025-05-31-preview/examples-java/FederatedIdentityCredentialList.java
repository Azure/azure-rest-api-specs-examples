
/**
 * Samples for FederatedIdentityCredentials List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/FederatedIdentityCredentialList.json
     */
    /**
     * Sample code: FederatedIdentityCredentialList.
     * 
     * @param manager Entry point to MsiManager.
     */
    public static void federatedIdentityCredentialList(com.azure.resourcemanager.msi.MsiManager manager) {
        manager.serviceClient().getFederatedIdentityCredentials().list("rgName", "resourceName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
