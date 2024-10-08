
import com.azure.resourcemanager.trustedsigning.models.RevokeCertificate;
import java.time.OffsetDateTime;

/**
 * Samples for CertificateProfiles RevokeCertificate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-05-preview/CertificateProfiles_RevokeCertificate.json
     */
    /**
     * Sample code: Revoke a certificate under a certificate profile.
     * 
     * @param manager Entry point to TrustedSigningManager.
     */
    public static void revokeACertificateUnderACertificateProfile(
        com.azure.resourcemanager.trustedsigning.TrustedSigningManager manager) {
        manager.certificateProfiles().revokeCertificateWithResponse("MyResourceGroup", "MyAccount", "profileA",
            new RevokeCertificate().withSerialNumber("xxxxxxxxxxxxxxxxxx")
                .withThumbprint("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
                .withEffectiveAt(OffsetDateTime.parse("2023-11-12T23:40:25+00:00")).withReason("KeyCompromised")
                .withRemarks("test"),
            com.azure.core.util.Context.NONE);
    }
}
