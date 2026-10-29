
import com.azure.resourcemanager.cdn.models.AfdOriginUpdateParameters;
import com.azure.resourcemanager.cdn.models.CertificateNameCheckValidationMode;
import com.azure.resourcemanager.cdn.models.EnabledState;

/**
 * Samples for AfdOrigins Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDOrigins_Update_OriginHostname.json
     */
    /**
     * Sample code: AFDOrigins_Update_OriginHostname.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDOriginsUpdateOriginHostname(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdOrigins().update("RG", "profile1", "origingroup1", "origin1",
            new AfdOriginUpdateParameters().withHostname("host1.blob.core.windows.net").withHttpPort(80)
                .withHttpsPort(443).withOriginHostHeader("host1.foo.com").withEnabledState(EnabledState.ENABLED)
                .withEnforceCertificateNameCheck(true).withCertificateNameCheckValidationMode(
                    CertificateNameCheckValidationMode.ORIGIN_HOSTNAME),
            com.azure.core.util.Context.NONE);
    }
}
