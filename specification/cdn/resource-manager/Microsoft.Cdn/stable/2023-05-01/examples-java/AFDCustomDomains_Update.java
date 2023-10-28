import com.azure.resourcemanager.cdn.models.AfdCertificateType;
import com.azure.resourcemanager.cdn.models.AfdDomainHttpsParameters;
import com.azure.resourcemanager.cdn.models.AfdDomainUpdateParameters;
import com.azure.resourcemanager.cdn.models.AfdMinimumTlsVersion;
import com.azure.resourcemanager.cdn.models.ResourceReference;

/** Samples for AfdCustomDomains Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/AFDCustomDomains_Update.json
     */
    /**
     * Sample code: AFDCustomDomains_Update.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDCustomDomainsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdCustomDomains()
            .update(
                "RG",
                "profile1",
                "domain1",
                new AfdDomainUpdateParameters()
                    .withTlsSettings(
                        new AfdDomainHttpsParameters()
                            .withCertificateType(AfdCertificateType.CUSTOMER_CERTIFICATE)
                            .withMinimumTlsVersion(AfdMinimumTlsVersion.TLS12))
                    .withAzureDnsZone(new ResourceReference().withId("")),
                com.azure.core.util.Context.NONE);
    }
}
