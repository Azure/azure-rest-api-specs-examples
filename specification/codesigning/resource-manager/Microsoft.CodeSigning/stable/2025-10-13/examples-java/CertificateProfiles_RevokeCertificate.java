
import com.azure.resourcemanager.artifactsigning.models.RevokeCertificate;
import java.time.OffsetDateTime;

/**
 * Samples for CertificateProfiles RevokeCertificate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-13/CertificateProfiles_RevokeCertificate.json
     */
    /**
     * Sample code: Revoke a certificate under a certificate profile.
     * 
     * @param manager Entry point to ArtifactSigningManager.
     */
    public static void revokeACertificateUnderACertificateProfile(
        com.azure.resourcemanager.artifactsigning.ArtifactSigningManager manager) {
        manager.certificateProfiles().revokeCertificateWithResponse("MyResourceGroup", "MyAccount", "profileA",
            new RevokeCertificate().withSerialNumber("xxxxxxxxxxxxxxxxxx")
                .withThumbprint("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
                .withEffectiveAt(OffsetDateTime.parse("2023-11-12T23:40:25+00:00")).withReason("KeyCompromised")
                .withRemarks("test"),
            com.azure.core.util.Context.NONE);
    }
}
