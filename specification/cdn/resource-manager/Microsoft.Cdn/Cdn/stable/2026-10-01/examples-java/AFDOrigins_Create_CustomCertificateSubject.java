
import com.azure.resourcemanager.cdn.fluent.models.AfdOriginInner;
import com.azure.resourcemanager.cdn.models.CertificateNameCheckValidationMode;
import com.azure.resourcemanager.cdn.models.EnabledState;
import java.util.Arrays;

/**
 * Samples for AfdOrigins Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDOrigins_Create_CustomCertificateSubject.json
     */
    /**
     * Sample code: AFDOrigins_Create_CustomCertificateSubject.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDOriginsCreateCustomCertificateSubject(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdOrigins().create("RG", "profile1", "origingroup1", "origin1",
            new AfdOriginInner().withHostname("host1.blob.core.windows.net").withHttpPort(80).withHttpsPort(443)
                .withOriginHostHeader("host1.foo.com").withEnabledState(EnabledState.ENABLED)
                .withEnforceCertificateNameCheck(true)
                .withCertificateNameCheckValidationMode(CertificateNameCheckValidationMode.CUSTOM_CERTIFICATE_SUBJECT)
                .withCustomCertificateSubjects(Arrays.asList("host1.foo.com", "*.foo.com")),
            com.azure.core.util.Context.NONE);
    }
}
