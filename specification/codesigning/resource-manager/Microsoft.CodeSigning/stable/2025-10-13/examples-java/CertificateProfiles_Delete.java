
/**
 * Samples for CertificateProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-13/CertificateProfiles_Delete.json
     */
    /**
     * Sample code: Delete a certificate profile.
     * 
     * @param manager Entry point to ArtifactSigningManager.
     */
    public static void
        deleteACertificateProfile(com.azure.resourcemanager.artifactsigning.ArtifactSigningManager manager) {
        manager.certificateProfiles().delete("MyResourceGroup", "MyAccount", "profileA",
            com.azure.core.util.Context.NONE);
    }
}
