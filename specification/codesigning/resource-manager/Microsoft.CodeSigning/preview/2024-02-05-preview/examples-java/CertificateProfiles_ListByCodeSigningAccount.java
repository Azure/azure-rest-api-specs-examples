
/**
 * Samples for CertificateProfiles ListByCodeSigningAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-05-preview/CertificateProfiles_ListByCodeSigningAccount.json
     */
    /**
     * Sample code: List certificate profiles under a trusted signing account.
     * 
     * @param manager Entry point to TrustedSigningManager.
     */
    public static void listCertificateProfilesUnderATrustedSigningAccount(
        com.azure.resourcemanager.trustedsigning.TrustedSigningManager manager) {
        manager.certificateProfiles().listByCodeSigningAccount("MyResourceGroup", "MyAccount",
            com.azure.core.util.Context.NONE);
    }
}
