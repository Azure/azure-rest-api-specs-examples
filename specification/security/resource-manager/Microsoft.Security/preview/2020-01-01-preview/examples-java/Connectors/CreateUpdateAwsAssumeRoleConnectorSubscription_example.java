
import com.azure.resourcemanager.security.models.AutoProvision;
import com.azure.resourcemanager.security.models.AwAssumeRoleAuthenticationDetailsProperties;
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
     * CreateUpdateAwsAssumeRoleConnectorSubscription_example.json
     */
    /**
     * Sample code: AwsAssumeRole - Create a cloud account connector for a subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void awsAssumeRoleCreateACloudAccountConnectorForASubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.connectors().define("aws_dev2")
            .withHybridComputeSettings(new HybridComputeSettingsProperties().withAutoProvision(AutoProvision.ON)
                .withResourceGroupName("AwsConnectorRG").withRegion("West US 2")
                .withProxyServer(new ProxyServerProperties().withIp("167.220.197.140").withPort("34"))
                .withServicePrincipal(new ServicePrincipalProperties()
                    .withApplicationId("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1").withSecret("fakeTokenPlaceholder")))
            .withAuthenticationDetails(new AwAssumeRoleAuthenticationDetailsProperties()
                .withAwsAssumeRoleArn("arn:aws:iam::81231569658:role/AscConnector")
                .withAwsExternalId("20ff7fc3-e762-44dd-bd96-b71116dcdc23"))
            .create();
    }
}
