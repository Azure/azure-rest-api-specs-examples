
import com.azure.resourcemanager.deviceregistry.models.CertificateConfiguration;
import com.azure.resourcemanager.deviceregistry.models.LeafCertificateConfiguration;
import com.azure.resourcemanager.deviceregistry.models.Policy;
import com.azure.resourcemanager.deviceregistry.models.PolicyUpdateProperties;

/**
 * Samples for Policies Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Update_Policies.json
     */
    /**
     * Sample code: Update_Policies.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void updatePolicies(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        Policy resource = manager.policies()
            .getWithResponse("rgdeviceregistry", "mynamespace", "mypolicy", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update()
            .withProperties(new PolicyUpdateProperties().withCertificate(new CertificateConfiguration()
                .withLeafCertificateConfiguration(new LeafCertificateConfiguration().withValidityPeriodInDays(10))))
            .apply();
    }
}
