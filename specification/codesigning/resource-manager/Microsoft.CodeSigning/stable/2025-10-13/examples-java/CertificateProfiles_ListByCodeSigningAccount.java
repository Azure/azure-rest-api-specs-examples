
/**
 * Samples for CertificateProfiles ListByCodeSigningAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-13/CertificateProfiles_ListByCodeSigningAccount.json
     */
    /**
     * Sample code: List certificate profiles under an artifact signing account.
     * 
     * @param manager Entry point to ArtifactSigningManager.
     */
    public static void listCertificateProfilesUnderAnArtifactSigningAccount(
        com.azure.resourcemanager.artifactsigning.ArtifactSigningManager manager) {
        manager.certificateProfiles().listByCodeSigningAccount("MyResourceGroup", "MyAccount",
            com.azure.core.util.Context.NONE);
    }
}
