
import com.azure.resourcemanager.deviceregistry.models.BringYourOwnRoot;
import com.azure.resourcemanager.deviceregistry.models.CertificateAuthorityConfiguration;
import com.azure.resourcemanager.deviceregistry.models.CertificateConfiguration;
import com.azure.resourcemanager.deviceregistry.models.LeafCertificateConfiguration;
import com.azure.resourcemanager.deviceregistry.models.PolicyProperties;
import com.azure.resourcemanager.deviceregistry.models.SupportedKeyType;

/**
 * Samples for Policies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/CreateOrReplace_Policies.json
     */
    /**
     * Sample code: CreateOrReplace_Policies.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void createOrReplacePolicies(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.policies().define("mypolicy").withExistingNamespace("rgdeviceregistry", "mynamespace")
            .withProperties(new PolicyProperties().withCertificate(new CertificateConfiguration()
                .withCertificateAuthorityConfiguration(new CertificateAuthorityConfiguration()
                    .withKeyType(SupportedKeyType.ECC).withBringYourOwnRoot(new BringYourOwnRoot().withEnabled(true)))
                .withLeafCertificateConfiguration(new LeafCertificateConfiguration().withValidityPeriodInDays(10))))
            .create();
    }
}
