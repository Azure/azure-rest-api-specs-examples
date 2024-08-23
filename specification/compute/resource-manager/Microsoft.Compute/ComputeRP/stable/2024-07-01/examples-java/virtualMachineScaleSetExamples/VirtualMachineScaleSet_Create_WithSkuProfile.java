
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetInner;
import com.azure.resourcemanager.compute.models.AllocationStrategy;
import com.azure.resourcemanager.compute.models.ApiEntityReference;
import com.azure.resourcemanager.compute.models.BillingProfile;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.OrchestrationMode;
import com.azure.resourcemanager.compute.models.PriorityMixPolicy;
import com.azure.resourcemanager.compute.models.Sku;
import com.azure.resourcemanager.compute.models.SkuProfile;
import com.azure.resourcemanager.compute.models.SkuProfileVMSize;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.VirtualMachineEvictionPolicyTypes;
import com.azure.resourcemanager.compute.models.VirtualMachinePriorityTypes;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetIpConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetManagedDiskParameters;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetNetworkConfiguration;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetNetworkProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetOSDisk;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetOSProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetStorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetVMProfile;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithSkuProfile.json
     */
    /**
     * Sample code: Create a scale set with sku profile.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAScaleSetWithSkuProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().createOrUpdate("myResourceGroup",
            "{vmss-name}",
            new VirtualMachineScaleSetInner().withLocation("westus")
                .withSku(new Sku().withName("Mix").withCapacity(10L))
                .withVirtualMachineProfile(new VirtualMachineScaleSetVMProfile()
                    .withOsProfile(new VirtualMachineScaleSetOSProfile().withComputerNamePrefix("{vmss-name}")
                        .withAdminUsername("{your-username}").withAdminPassword("fakeTokenPlaceholder"))
                    .withStorageProfile(new VirtualMachineScaleSetStorageProfile()
                        .withImageReference(new ImageReference().withPublisher("MicrosoftWindowsServer")
                            .withOffer("WindowsServer").withSku("2016-Datacenter").withVersion("latest"))
                        .withOsDisk(new VirtualMachineScaleSetOSDisk().withCaching(CachingTypes.READ_WRITE)
                            .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                            .withManagedDisk(new VirtualMachineScaleSetManagedDiskParameters()
                                .withStorageAccountType(StorageAccountTypes.STANDARD_LRS))))
                    .withNetworkProfile(new VirtualMachineScaleSetNetworkProfile().withNetworkInterfaceConfigurations(
                        Arrays.asList(new VirtualMachineScaleSetNetworkConfiguration().withName("{vmss-name}")
                            .withPrimary(true)
                            .withIpConfigurations(Arrays.asList(new VirtualMachineScaleSetIpConfiguration()
                                .withName("{vmss-name}")
                                .withSubnet(new ApiEntityReference().withId(
                                    "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"))))
                            .withEnableIpForwarding(true))))
                    .withPriority(VirtualMachinePriorityTypes.SPOT)
                    .withEvictionPolicy(VirtualMachineEvictionPolicyTypes.DEALLOCATE)
                    .withBillingProfile(new BillingProfile().withMaxPrice(-1.0D)))
                .withSinglePlacementGroup(false).withOrchestrationMode(OrchestrationMode.FLEXIBLE)
                .withPriorityMixPolicy(
                    new PriorityMixPolicy().withBaseRegularPriorityCount(4).withRegularPriorityPercentageAboveBase(50))
                .withSkuProfile(new SkuProfile()
                    .withVmSizes(Arrays.asList(new SkuProfileVMSize().withName("Standard_D8s_v5"),
                        new SkuProfileVMSize().withName("Standard_E16s_v5"),
                        new SkuProfileVMSize().withName("Standard_D2s_v5")))
                    .withAllocationStrategy(AllocationStrategy.CAPACITY_OPTIMIZED)),
            null, null, com.azure.core.util.Context.NONE);
    }
}
