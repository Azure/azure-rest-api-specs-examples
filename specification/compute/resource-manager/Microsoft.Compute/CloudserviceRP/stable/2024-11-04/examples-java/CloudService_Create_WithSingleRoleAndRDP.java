
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.compute.fluent.models.CloudServiceInner;
import com.azure.resourcemanager.compute.models.CloudServiceExtensionProfile;
import com.azure.resourcemanager.compute.models.CloudServiceExtensionProperties;
import com.azure.resourcemanager.compute.models.CloudServiceNetworkProfile;
import com.azure.resourcemanager.compute.models.CloudServiceProperties;
import com.azure.resourcemanager.compute.models.CloudServiceRoleProfile;
import com.azure.resourcemanager.compute.models.CloudServiceRoleProfileProperties;
import com.azure.resourcemanager.compute.models.CloudServiceRoleSku;
import com.azure.resourcemanager.compute.models.CloudServiceUpgradeMode;
import com.azure.resourcemanager.compute.models.Extension;
import com.azure.resourcemanager.compute.models.LoadBalancerConfiguration;
import com.azure.resourcemanager.compute.models.LoadBalancerConfigurationProperties;
import com.azure.resourcemanager.compute.models.LoadBalancerFrontendIpConfiguration;
import com.azure.resourcemanager.compute.models.LoadBalancerFrontendIpConfigurationProperties;
import java.util.Arrays;

/**
 * Samples for CloudServices CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/
     * CloudService_Create_WithSingleRoleAndRDP.json
     */
    /**
     * Sample code: Create New Cloud Service with Single Role and RDP Extension.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createNewCloudServiceWithSingleRoleAndRDPExtension(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServices().createOrUpdate("ConstosoRG", "{cs-name}",
            new CloudServiceInner().withLocation("westus").withProperties(new CloudServiceProperties()
                .withPackageUrl("{PackageUrl}").withConfiguration("{ServiceConfiguration}")
                .withUpgradeMode(CloudServiceUpgradeMode.AUTO)
                .withRoleProfile(new CloudServiceRoleProfile().withRoles(
                    Arrays.asList(new CloudServiceRoleProfileProperties().withName("ContosoFrontend").withSku(
                        new CloudServiceRoleSku().withName("Standard_D1_v2").withTier("Standard").withCapacity(1L)))))
                .withNetworkProfile(new CloudServiceNetworkProfile()
                    .withLoadBalancerConfigurations(Arrays.asList(new LoadBalancerConfiguration().withName("contosolb")
                        .withProperties(new LoadBalancerConfigurationProperties().withFrontendIpConfigurations(
                            Arrays.asList(new LoadBalancerFrontendIpConfiguration().withName("contosofe")
                                .withProperties(new LoadBalancerFrontendIpConfigurationProperties()
                                    .withPublicIpAddress(new SubResource().withId(
                                        "/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Network/publicIPAddresses/contosopublicip")))))))))
                .withExtensionProfile(new CloudServiceExtensionProfile().withExtensions(Arrays.asList(new Extension()
                    .withName("RDPExtension")
                    .withProperties(new CloudServiceExtensionProperties()
                        .withPublisher("Microsoft.Windows.Azure.Extensions").withType("RDP")
                        .withTypeHandlerVersion("1.2").withAutoUpgradeMinorVersion(false)
                        .withSettings(
                            "<PublicConfig><UserName>UserAzure</UserName><Expiration>10/22/2021 15:05:45</Expiration></PublicConfig>")
                        .withProtectedSettings("<PrivateConfig><Password>{password}</Password></PrivateConfig>")))))),
            com.azure.core.util.Context.NONE);
    }
}
