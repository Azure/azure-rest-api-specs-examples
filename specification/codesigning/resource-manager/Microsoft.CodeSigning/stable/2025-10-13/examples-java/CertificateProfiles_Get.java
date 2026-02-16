
/**
 * Samples for CertificateProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-13/CertificateProfiles_Get.json
     */
    /**
     * Sample code: Get details of a certificate profile.
     * 
     * @param manager Entry point to ArtifactSigningManager.
     */
    public static void
        getDetailsOfACertificateProfile(com.azure.resourcemanager.artifactsigning.ArtifactSigningManager manager) {
        manager.certificateProfiles().getWithResponse("MyResourceGroup", "MyAccount", "profileA",
            com.azure.core.util.Context.NONE);
    }
}
