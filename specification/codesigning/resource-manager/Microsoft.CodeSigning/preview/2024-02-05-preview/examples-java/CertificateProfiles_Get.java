
/**
 * Samples for CertificateProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-05-preview/CertificateProfiles_Get.json
     */
    /**
     * Sample code: Get details of a certificate profile.
     * 
     * @param manager Entry point to TrustedSigningManager.
     */
    public static void
        getDetailsOfACertificateProfile(com.azure.resourcemanager.trustedsigning.TrustedSigningManager manager) {
        manager.certificateProfiles().getWithResponse("MyResourceGroup", "MyAccount", "profileA",
            com.azure.core.util.Context.NONE);
    }
}
