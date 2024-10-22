
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.computefleet.models.ApiEntityReference;
import com.azure.resourcemanager.computefleet.models.BaseVirtualMachineProfile;
import com.azure.resourcemanager.computefleet.models.CachingTypes;
import com.azure.resourcemanager.computefleet.models.ComputeProfile;
import com.azure.resourcemanager.computefleet.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.computefleet.models.EvictionPolicy;
import com.azure.resourcemanager.computefleet.models.FleetProperties;
import com.azure.resourcemanager.computefleet.models.ImageReference;
import com.azure.resourcemanager.computefleet.models.LinuxConfiguration;
import com.azure.resourcemanager.computefleet.models.NetworkApiVersion;
import com.azure.resourcemanager.computefleet.models.OperatingSystemTypes;
import com.azure.resourcemanager.computefleet.models.RegularPriorityAllocationStrategy;
import com.azure.resourcemanager.computefleet.models.RegularPriorityProfile;
import com.azure.resourcemanager.computefleet.models.SpotAllocationStrategy;
import com.azure.resourcemanager.computefleet.models.SpotPriorityProfile;
import com.azure.resourcemanager.computefleet.models.StorageAccountTypes;
import com.azure.resourcemanager.computefleet.models.VirtualMachineScaleSetIPConfiguration;
import com.azure.resourcemanager.computefleet.models.VirtualMachineScaleSetIPConfigurationProperties;
import com.azure.resourcemanager.computefleet.models.VirtualMachineScaleSetManagedDiskParameters;
import com.azure.resourcemanager.computefleet.models.VirtualMachineScaleSetNetworkConfiguration;
import com.azure.resourcemanager.computefleet.models.VirtualMachineScaleSetNetworkConfigurationProperties;
import com.azure.resourcemanager.computefleet.models.VirtualMachineScaleSetNetworkProfile;
import com.azure.resourcemanager.computefleet.models.VirtualMachineScaleSetOSDisk;
import com.azure.resourcemanager.computefleet.models.VirtualMachineScaleSetOSProfile;
import com.azure.resourcemanager.computefleet.models.VirtualMachineScaleSetStorageProfile;
import com.azure.resourcemanager.computefleet.models.VmSizeProfile;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Fleets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Fleets_CreateOrUpdate_MinimumSet.json
     */
    /**
     * Sample code: Fleets_CreateOrUpdate_MinimumSet.
     * 
     * @param manager Entry point to ComputeFleetManager.
     */
    public static void
        fleetsCreateOrUpdateMinimumSet(com.azure.resourcemanager.computefleet.ComputeFleetManager manager) {
        manager.fleets().define("testFleet").withRegion("eastus2euap").withExistingResourceGroup("rgazurefleet")
            .withTags(mapOf("key", "fakeTokenPlaceholder"))
            .withProperties(new FleetProperties()
                .withSpotPriorityProfile(new SpotPriorityProfile().withCapacity(2).withMinCapacity(1)
                    .withEvictionPolicy(EvictionPolicy.DELETE)
                    .withAllocationStrategy(SpotAllocationStrategy.PRICE_CAPACITY_OPTIMIZED).withMaintain(true))
                .withRegularPriorityProfile(new RegularPriorityProfile().withCapacity(2).withMinCapacity(1)
                    .withAllocationStrategy(RegularPriorityAllocationStrategy.LOWEST_PRICE))
                .withVmSizesProfile(Arrays.asList(new VmSizeProfile().withName("Standard_D2s_v3"),
                    new VmSizeProfile().withName("Standard_D4s_v3"), new VmSizeProfile().withName("Standard_E2s_v3")))
                .withComputeProfile(new ComputeProfile()
                    .withBaseVirtualMachineProfile(new BaseVirtualMachineProfile()
                        .withOsProfile(new VirtualMachineScaleSetOSProfile().withComputerNamePrefix("prefix")
                            .withAdminUsername("azureuser").withAdminPassword("fakeTokenPlaceholder")
                            .withLinuxConfiguration(new LinuxConfiguration().withDisablePasswordAuthentication(false)))
                        .withStorageProfile(new VirtualMachineScaleSetStorageProfile()
                            .withImageReference(new ImageReference()
                                .withPublisher("canonical").withOffer("0001-com-ubuntu-server-focal")
                                .withSku("20_04-lts-gen2").withVersion("latest"))
                            .withOsDisk(new VirtualMachineScaleSetOSDisk().withCaching(CachingTypes.READ_WRITE)
                                .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                                .withOsType(OperatingSystemTypes.LINUX)
                                .withManagedDisk(new VirtualMachineScaleSetManagedDiskParameters()
                                    .withStorageAccountType(StorageAccountTypes.STANDARD_LRS))))
                        .withNetworkProfile(new VirtualMachineScaleSetNetworkProfile()
                            .withNetworkInterfaceConfigurations(
                                Arrays.asList(new VirtualMachineScaleSetNetworkConfiguration().withName("vmNameTest")
                                    .withProperties(new VirtualMachineScaleSetNetworkConfigurationProperties()
                                        .withPrimary(true).withEnableAcceleratedNetworking(false)
                                        .withIpConfigurations(Arrays.asList(new VirtualMachineScaleSetIPConfiguration()
                                            .withName("vmNameTest")
                                            .withProperties(new VirtualMachineScaleSetIPConfigurationProperties()
                                                .withSubnet(new ApiEntityReference().withId(
                                                    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"))
                                                .withPrimary(true).withLoadBalancerBackendAddressPools(
                                                    Arrays.asList(new SubResource().withId(
                                                        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"))))))
                                        .withEnableIPForwarding(true))))
                            .withNetworkApiVersion(NetworkApiVersion.fromString("2022-07-01"))))
                    .withComputeApiVersion("2023-09-01").withPlatformFaultDomainCount(1)))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
