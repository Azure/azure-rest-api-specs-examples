
import com.azure.resourcemanager.security.models.AutoProvision;
import com.azure.resourcemanager.security.models.GcpCredentialsDetailsProperties;
import com.azure.resourcemanager.security.models.HybridComputeSettingsProperties;

/**
 * Samples for Connectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/
     * CreateUpdateGcpCredentialsConnectorSubscription_example.json
     */
    /**
     * Sample code: gcpCredentials - Create a cloud account connector for a subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void gcpCredentialsCreateACloudAccountConnectorForASubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.connectors().define("gcp_dev")
            .withHybridComputeSettings(new HybridComputeSettingsProperties().withAutoProvision(AutoProvision.OFF))
            .withAuthenticationDetails(new GcpCredentialsDetailsProperties().withOrganizationId("AscDemoOrg")
                .withType("service_account").withProjectId("asc-project-1234").withPrivateKeyId("fakeTokenPlaceholder")
                .withPrivateKey("fakeTokenPlaceholder")
                .withClientEmail("asc-135@asc-project-1234.iam.gserviceaccount.com")
                .withClientId("105889053725632919854").withAuthUri("https://accounts.google.com/o/oauth2/auth")
                .withTokenUri("fakeTokenPlaceholder")
                .withAuthProviderX509CertUrl("https://www.googleapis.com/oauth2/v1/certs").withClientX509CertUrl(
                    "https://www.googleapis.com/robot/v1/metadata/x509/asc-135%40asc-project-1234.iam.gserviceaccount.com"))
            .create();
    }
}
