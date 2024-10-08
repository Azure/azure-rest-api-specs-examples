
import com.azure.resourcemanager.trustedsigning.models.CertificateProfileProperties;
import com.azure.resourcemanager.trustedsigning.models.ProfileType;

/**
 * Samples for CertificateProfiles Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-05-preview/CertificateProfiles_Create.json
     */
    /**
     * Sample code: Create a certificate profile.
     * 
     * @param manager Entry point to TrustedSigningManager.
     */
    public static void
        createACertificateProfile(com.azure.resourcemanager.trustedsigning.TrustedSigningManager manager) {
        manager.certificateProfiles().define("profileA").withExistingCodeSigningAccount("MyResourceGroup", "MyAccount")
            .withProperties(new CertificateProfileProperties().withProfileType(ProfileType.PUBLIC_TRUST)
                .withIncludeStreetAddress(false).withIncludePostalCode(true)
                .withIdentityValidationId("00000000-1234-5678-3333-444444444444"))
            .create();
    }
}
