
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.DomainNameLabelScopeTypes;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.IpVersion;
import com.azure.resourcemanager.compute.models.NetworkInterfaceAuxiliaryMode;
import com.azure.resourcemanager.compute.models.NetworkInterfaceAuxiliarySku;
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
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetPublicIpAddressConfigurationDnsSettings;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetStorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMProfile;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_FromWithNetworkInterfaceWithDnsSettings.json
     */
    /**
     * Sample code: Create a scale set with Network Interfaces with public ip address dns settings.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAScaleSetWithNetworkInterfacesWithPublicIpAddressDnsSettings(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets()
            .createOrUpdate("myResourceGroup", "{vmss-name}", new VirtualMachineScaleSetInner().withLocation("westus")
                .withSku(new Sku().withName("Standard_D1_v2").withTier("Standard").withCapacity(3L))
                .withUpgradePolicy(new UpgradePolicy().withMode(UpgradeMode.MANUAL))
                .withVirtualMachineProfile(new VirtualMachineScaleSetVMProfile()
                    .withOsProfile(new VirtualMachineScaleSetOSProfile().withComputerNamePrefix("{vmss-name}")
                        .withAdminUsername("{your-username}").withAdminPassword("fakeTokenPlaceholder"))
                    .withStorageProfile(new VirtualMachineScaleSetStorageProfile()
                        .withImageReference(new ImageReference().withId(
                            "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"))
                        .withOsDisk(new VirtualMachineScaleSetOSDisk().withCaching(CachingTypes.READ_WRITE)
                            .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                            .withManagedDisk(new VirtualMachineScaleSetManagedDiskParameters()
                                .withStorageAccountType(StorageAccountTypes.STANDARD_LRS))))
                    .withNetworkProfile(new VirtualMachineScaleSetNetworkProfile().withNetworkInterfaceConfigurations(
                        Arrays.asList(new VirtualMachineScaleSetNetworkConfiguration().withName("{nicConfig1-name}")
                            .withPrimary(true).withEnableAcceleratedNetworking(true).withDisableTcpStateTracking(true)
                            .withIpConfigurations(Arrays.asList(new VirtualMachineScaleSetIpConfiguration()
                                .withName("{vmss-name}")
                                .withSubnet(new ApiEntityReference().withId(
                                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"))))
                            .withEnableIpForwarding(true)
                            .withAuxiliaryMode(NetworkInterfaceAuxiliaryMode.ACCELERATED_CONNECTIONS)
                            .withAuxiliarySku(NetworkInterfaceAuxiliarySku.A1),
                            new VirtualMachineScaleSetNetworkConfiguration().withName("{nicConfig2-name}")
                                .withPrimary(false).withEnableAcceleratedNetworking(false)
                                .withDisableTcpStateTracking(false)
                                .withIpConfigurations(Arrays.asList(new VirtualMachineScaleSetIpConfiguration()
                                    .withName("{nicConfig2-name}")
                                    .withSubnet(new ApiEntityReference().withId(
                                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-fpga-subnet-name2}"))
                                    .withPrimary(true)
                                    .withPublicIpAddressConfiguration(
                                        new VirtualMachineScaleSetPublicIpAddressConfiguration().withName("publicip")
                                            .withIdleTimeoutInMinutes(10).withDnsSettings(
                                                new VirtualMachineScaleSetPublicIpAddressConfigurationDnsSettings()
                                                    .withDomainNameLabel("vmsstestlabel01")
                                                    .withDomainNameLabelScope(DomainNameLabelScopeTypes.NO_REUSE)))
                                    .withPrivateIpAddressVersion(IpVersion.IPV4)))
                                .withEnableIpForwarding(false)))))
                .withOverprovision(true), null, null, com.azure.core.util.Context.NONE);
    }
}
