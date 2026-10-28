
import com.azure.resourcemanager.cdn.fluent.models.AfdOriginInner;
import com.azure.resourcemanager.cdn.models.CertificateNameCheckValidationMode;
import com.azure.resourcemanager.cdn.models.EnabledState;

/**
 * Samples for AfdOrigins Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDOrigins_Create_OriginHostname.json
     */
    /**
     * Sample code: AFDOrigins_Create_OriginHostname.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDOriginsCreateOriginHostname(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdOrigins().create("RG", "profile1", "origingroup1", "origin1",
            new AfdOriginInner().withHostname("host1.blob.core.windows.net").withHttpPort(80).withHttpsPort(443)
                .withOriginHostHeader("host1.foo.com").withEnabledState(EnabledState.ENABLED)
                .withEnforceCertificateNameCheck(true).withCertificateNameCheckValidationMode(
                    CertificateNameCheckValidationMode.ORIGIN_HOSTNAME),
            com.azure.core.util.Context.NONE);
    }
}
