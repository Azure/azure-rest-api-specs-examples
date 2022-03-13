Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.IpVersion;
import com.azure.resourcemanager.compute.models.Sku;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.UpgradeMode;
import com.azure.resourcemanager.compute.models.UpgradePolicy;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetIpConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetManagedDiskParameters;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetNetworkConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetNetworkProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetOSDisk;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetOSProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetPublicIpAddressConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetStorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMProfile;
import java.util.Arrays;

/** Samples for VirtualMachineScaleSets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateAScaleSetWithAnAzureLoadBalancer.json
     */
    /**
     * Sample code: Create a scale set with an azure load balancer.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAScaleSetWithAnAzureLoadBalancer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .createOrUpdate(
                "myResourceGroup",
                "{vmss-name}",
                new VirtualMachineScaleSetInner()
                    .withLocation("westus")
                    .withSku(new Sku().withName("Standard_D1_v2").withTier("Standard").withCapacity(3L))
                    .withUpgradePolicy(new UpgradePolicy().withMode(UpgradeMode.MANUAL))
                    .withVirtualMachineProfile(
                        new VirtualMachineScaleSetVMProfile()
                            .withOsProfile(
                                new VirtualMachineScaleSetOSProfile()
                                    .withComputerNamePrefix("{vmss-name}")
                                    .withAdminUsername("{your-username}")
                                    .withAdminPassword("{your-password}"))
                            .withStorageProfile(
                                new VirtualMachineScaleSetStorageProfile()
                                    .withImageReference(
                                        new ImageReference()
                                            .withPublisher("MicrosoftWindowsServer")
                                            .withOffer("WindowsServer")
                                            .withSku("2016-Datacenter")
                                            .withVersion("latest"))
                                    .withOsDisk(
                                        new VirtualMachineScaleSetOSDisk()
                                            .withCaching(CachingTypes.READ_WRITE)
                                            .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                                            .withManagedDisk(
                                                new VirtualMachineScaleSetManagedDiskParameters()
                                                    .withStorageAccountType(StorageAccountTypes.STANDARD_LRS))))
                            .withNetworkProfile(
                                new VirtualMachineScaleSetNetworkProfile()
                                    .withNetworkInterfaceConfigurations(
                                        Arrays
                                            .asList(
                                                new VirtualMachineScaleSetNetworkConfiguration()
                                                    .withName("{vmss-name}")
                                                    .withPrimary(true)
                                                    .withIpConfigurations(
                                                        Arrays
                                                            .asList(
                                                                new VirtualMachineScaleSetIpConfiguration()
                                                                    .withName("{vmss-name}")
                                                                    .withSubnet(
                                                                        new ApiEntityReference()
                                                                            .withId(
                                                                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"))
                                                                    .withPublicIpAddressConfiguration(
                                                                        new VirtualMachineScaleSetPublicIpAddressConfiguration()
                                                                            .withName("{vmss-name}")
                                                                            .withPublicIpAddressVersion(IpVersion.IPV4))
                                                                    .withLoadBalancerBackendAddressPools(
                                                                        Arrays
                                                                            .asList(
                                                                                new SubResource()
                                                                                    .withId(
                                                                                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/loadBalancers/{existing-load-balancer-name}/backendAddressPools/{existing-backend-address-pool-name}")))
                                                                    .withLoadBalancerInboundNatPools(
                                                                        Arrays
                                                                            .asList(
                                                                                new SubResource()
                                                                                    .withId(
                                                                                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/loadBalancers/{existing-load-balancer-name}/inboundNatPools/{existing-nat-pool-name}")))))
                                                    .withEnableIpForwarding(true)))))
                    .withOverprovision(true),
                Context.NONE);
    }
}
```
