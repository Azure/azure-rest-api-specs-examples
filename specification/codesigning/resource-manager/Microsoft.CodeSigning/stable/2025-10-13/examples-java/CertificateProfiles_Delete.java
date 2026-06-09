
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
     * @param manager Entry point to TrustedSigningManager.
     */
    public static void
        deleteACertificateProfile(com.azure.resourcemanager.trustedsigning.TrustedSigningManager manager) {
        manager.certificateProfiles().delete("MyResourceGroup", "MyAccount", "profileA",
            com.azure.core.util.Context.NONE);
    }
}
