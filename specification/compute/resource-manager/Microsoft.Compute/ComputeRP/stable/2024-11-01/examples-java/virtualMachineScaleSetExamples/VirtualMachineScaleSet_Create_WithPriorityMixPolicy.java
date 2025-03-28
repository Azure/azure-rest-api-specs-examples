
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetInner;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.NetworkApiVersion;
import com.azure.resourcemanager.compute.models.OrchestrationMode;
import com.azure.resourcemanager.compute.models.PriorityMixPolicy;
import com.azure.resourcemanager.compute.models.Sku;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.VirtualMachinePriorityTypes;
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

/**
 * Samples for VirtualMachineScaleSets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithPriorityMixPolicy.json
     */
    /**
     * Sample code: Create a scale set with priority mix policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAScaleSetWithPriorityMixPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets()
            .createOrUpdate("myResourceGroup", "{vmss-name}",
                new VirtualMachineScaleSetInner().withLocation("westus")
                    .withSku(new Sku().withName(
                        "Standard_A8m_v2").withTier(
                            "Standard")
                        .withCapacity(2L))
                    .withVirtualMachineProfile(new VirtualMachineScaleSetVMProfile()
                        .withOsProfile(new VirtualMachineScaleSetOSProfile().withComputerNamePrefix("{vmss-name}")
                            .withAdminUsername("{your-username}"))
                        .withStorageProfile(new VirtualMachineScaleSetStorageProfile()
                            .withImageReference(new ImageReference()
                                .withPublisher("Canonical").withOffer("0001-com-ubuntu-server-focal").withSku(
                                    "20_04-lts-gen2")
                                .withVersion("latest"))
                            .withOsDisk(new VirtualMachineScaleSetOSDisk().withCaching(CachingTypes.READ_WRITE)
                                .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                                .withManagedDisk(
                                    new VirtualMachineScaleSetManagedDiskParameters()
                                        .withStorageAccountType(StorageAccountTypes.STANDARD_LRS))))
                        .withNetworkProfile(new VirtualMachineScaleSetNetworkProfile()
                            .withNetworkInterfaceConfigurations(
                                Arrays.asList(new VirtualMachineScaleSetNetworkConfiguration().withName("{vmss-name}")
                                    .withPrimary(true).withEnableAcceleratedNetworking(false)
                                    .withIpConfigurations(Arrays.asList(new VirtualMachineScaleSetIpConfiguration()
                                        .withName("{vmss-name}")
                                        .withSubnet(new ApiEntityReference().withId(
                                            "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"))
                                        .withPrimary(true)
                                        .withPublicIpAddressConfiguration(
                                            new VirtualMachineScaleSetPublicIpAddressConfiguration()
                                                .withName("{vmss-name}").withIdleTimeoutInMinutes(15))
                                        .withApplicationGatewayBackendAddressPools(Arrays.asList())
                                        .withLoadBalancerBackendAddressPools(Arrays.asList())))
                                    .withEnableIpForwarding(true)))
                            .withNetworkApiVersion(NetworkApiVersion.TWO_ZERO_TWO_ZERO_ONE_ONE_ZERO_ONE))
                        .withPriority(VirtualMachinePriorityTypes.SPOT))
                    .withPlatformFaultDomainCount(1).withOrchestrationMode(OrchestrationMode.FLEXIBLE)
                    .withPriorityMixPolicy(new PriorityMixPolicy().withBaseRegularPriorityCount(10)
                        .withRegularPriorityPercentageAboveBase(50)),
                null, null, com.azure.core.util.Context.NONE);
    }
}
