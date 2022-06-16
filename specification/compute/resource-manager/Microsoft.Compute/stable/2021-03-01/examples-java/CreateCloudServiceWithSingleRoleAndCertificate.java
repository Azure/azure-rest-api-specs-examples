import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.CloudServiceInner;
import com.azure.resourcemanager.compute.models.CloudServiceNetworkProfile;
import com.azure.resourcemanager.compute.models.CloudServiceOsProfile;
import com.azure.resourcemanager.compute.models.CloudServiceProperties;
import com.azure.resourcemanager.compute.models.CloudServiceRoleProfile;
import com.azure.resourcemanager.compute.models.CloudServiceRoleProfileProperties;
import com.azure.resourcemanager.compute.models.CloudServiceRoleSku;
import com.azure.resourcemanager.compute.models.CloudServiceUpgradeMode;
import com.azure.resourcemanager.compute.models.CloudServiceVaultCertificate;
import com.azure.resourcemanager.compute.models.CloudServiceVaultSecretGroup;
import com.azure.resourcemanager.compute.models.LoadBalancerConfiguration;
import com.azure.resourcemanager.compute.models.LoadBalancerConfigurationProperties;
import com.azure.resourcemanager.compute.models.LoadBalancerFrontendIpConfiguration;
import com.azure.resourcemanager.compute.models.LoadBalancerFrontendIpConfigurationProperties;
import java.util.Arrays;

/** Samples for CloudServices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/CreateCloudServiceWithSingleRoleAndCertificate.json
     */
    /**
     * Sample code: Create New Cloud Service with Single Role and Certificate from Key Vault.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNewCloudServiceWithSingleRoleAndCertificateFromKeyVault(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServices()
            .createOrUpdate(
                "ConstosoRG",
                "{cs-name}",
                new CloudServiceInner()
                    .withLocation("westus")
                    .withProperties(
                        new CloudServiceProperties()
                            .withPackageUrl("{PackageUrl}")
                            .withConfiguration("{ServiceConfiguration}")
                            .withUpgradeMode(CloudServiceUpgradeMode.AUTO)
                            .withRoleProfile(
                                new CloudServiceRoleProfile()
                                    .withRoles(
                                        Arrays
                                            .asList(
                                                new CloudServiceRoleProfileProperties()
                                                    .withName("ContosoFrontend")
                                                    .withSku(
                                                        new CloudServiceRoleSku()
                                                            .withName("Standard_D1_v2")
                                                            .withTier("Standard")
                                                            .withCapacity(1L)))))
                            .withOsProfile(
                                new CloudServiceOsProfile()
                                    .withSecrets(
                                        Arrays
                                            .asList(
                                                new CloudServiceVaultSecretGroup()
                                                    .withSourceVault(
                                                        new SubResource()
                                                            .withId(
                                                                "/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.KeyVault/vaults/{keyvault-name}"))
                                                    .withVaultCertificates(
                                                        Arrays
                                                            .asList(
                                                                new CloudServiceVaultCertificate()
                                                                    .withCertificateUrl(
                                                                        "https://{keyvault-name}.vault.azure.net:443/secrets/ContosoCertificate/{secret-id}"))))))
                            .withNetworkProfile(
                                new CloudServiceNetworkProfile()
                                    .withLoadBalancerConfigurations(
                                        Arrays
                                            .asList(
                                                new LoadBalancerConfiguration()
                                                    .withName("contosolb")
                                                    .withProperties(
                                                        new LoadBalancerConfigurationProperties()
                                                            .withFrontendIpConfigurations(
                                                                Arrays
                                                                    .asList(
                                                                        new LoadBalancerFrontendIpConfiguration()
                                                                            .withName("contosofe")
                                                                            .withProperties(
                                                                                new LoadBalancerFrontendIpConfigurationProperties()
                                                                                    .withPublicIpAddress(
                                                                                        new SubResource()
                                                                                            .withId(
                                                                                                "/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Network/publicIPAddresses/contosopublicip")))))))))),
                Context.NONE);
    }
}
