
/**
 * Samples for FederatedIdentityCredentials Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/FlexibleFederatedIdentityCredentialGet.json
     */
    /**
     * Sample code: FlexibleFederatedIdentityCredentialGet.
     * 
     * @param manager Entry point to MsiManager.
     */
    public static void flexibleFederatedIdentityCredentialGet(com.azure.resourcemanager.msi.MsiManager manager) {
        manager.serviceClient().getFederatedIdentityCredentials().getWithResponse("rgName", "resourceName",
            "ficResourceName", com.azure.core.util.Context.NONE);
    }
}
