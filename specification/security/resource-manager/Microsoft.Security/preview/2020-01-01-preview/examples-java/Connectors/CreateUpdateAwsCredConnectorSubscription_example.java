
import com.azure.resourcemanager.security.models.AutoProvision;
import com.azure.resourcemanager.security.models.AwsCredsAuthenticationDetailsProperties;
import com.azure.resourcemanager.security.models.HybridComputeSettingsProperties;
import com.azure.resourcemanager.security.models.ProxyServerProperties;
import com.azure.resourcemanager.security.models.ServicePrincipalProperties;

/**
 * Samples for Connectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/
     * CreateUpdateAwsCredConnectorSubscription_example.json
     */
    /**
     * Sample code: AwsCred - Create a cloud account connector for a subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void awsCredCreateACloudAccountConnectorForASubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.connectors().define("aws_dev1")
            .withHybridComputeSettings(new HybridComputeSettingsProperties().withAutoProvision(AutoProvision.ON)
                .withResourceGroupName("AwsConnectorRG").withRegion("West US 2")
                .withProxyServer(new ProxyServerProperties().withIp("167.220.197.140").withPort("34"))
                .withServicePrincipal(new ServicePrincipalProperties()
                    .withApplicationId("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1").withSecret("fakeTokenPlaceholder")))
            .withAuthenticationDetails(new AwsCredsAuthenticationDetailsProperties()
                .withAwsAccessKeyId("fakeTokenPlaceholder").withAwsSecretAccessKey("fakeTokenPlaceholder"))
            .create();
    }
}
