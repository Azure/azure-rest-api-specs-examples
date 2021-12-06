Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.CloudServiceInner;
import com.azure.resourcemanager.compute.models.CloudServiceNetworkProfile;
import com.azure.resourcemanager.compute.models.CloudServiceProperties;
import com.azure.resourcemanager.compute.models.CloudServiceRoleProfile;
import com.azure.resourcemanager.compute.models.CloudServiceRoleProfileProperties;
import com.azure.resourcemanager.compute.models.CloudServiceRoleSku;
import com.azure.resourcemanager.compute.models.CloudServiceUpgradeMode;
import com.azure.resourcemanager.compute.models.LoadBalancerConfiguration;
import com.azure.resourcemanager.compute.models.LoadBalancerConfigurationProperties;
import com.azure.resourcemanager.compute.models.LoadBalancerFrontendIpConfiguration;
import com.azure.resourcemanager.compute.models.LoadBalancerFrontendIpConfigurationProperties;
import java.util.Arrays;

/** Samples for CloudServices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/CreateCloudServiceWithMultiRole.json
     */
    /**
     * Sample code: Create New Cloud Service with Multiple Roles.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNewCloudServiceWithMultipleRoles(com.azure.resourcemanager.AzureResourceManager azure) {
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
                                                            .withCapacity(1L)),
                                                new CloudServiceRoleProfileProperties()
                                                    .withName("ContosoBackend")
                                                    .withSku(
                                                        new CloudServiceRoleSku()
                                                            .withName("Standard_D1_v2")
                                                            .withTier("Standard")
                                                            .withCapacity(1L)))))
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
```
