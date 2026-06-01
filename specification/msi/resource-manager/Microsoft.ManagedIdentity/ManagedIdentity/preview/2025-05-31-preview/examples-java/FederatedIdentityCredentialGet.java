
/**
 * Samples for FederatedIdentityCredentials Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/FederatedIdentityCredentialGet.json
     */
    /**
     * Sample code: FederatedIdentityCredentialGet.
     * 
     * @param manager Entry point to MsiManager.
     */
    public static void federatedIdentityCredentialGet(com.azure.resourcemanager.msi.MsiManager manager) {
        manager.serviceClient().getFederatedIdentityCredentials().getWithResponse("rgName", "resourceName",
            "ficResourceName", com.azure.core.util.Context.NONE);
    }
}
