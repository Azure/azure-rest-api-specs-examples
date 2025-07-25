
import com.azure.resourcemanager.networkcloud.models.DefaultGateway;
import com.azure.resourcemanager.networkcloud.models.ExtendedLocation;
import com.azure.resourcemanager.networkcloud.models.ImageRepositoryCredentials;
import com.azure.resourcemanager.networkcloud.models.NetworkAttachment;
import com.azure.resourcemanager.networkcloud.models.OsDisk;
import com.azure.resourcemanager.networkcloud.models.OsDiskCreateOption;
import com.azure.resourcemanager.networkcloud.models.OsDiskDeleteOption;
import com.azure.resourcemanager.networkcloud.models.SshPublicKey;
import com.azure.resourcemanager.networkcloud.models.StorageProfile;
import com.azure.resourcemanager.networkcloud.models.VirtualMachineBootMethod;
import com.azure.resourcemanager.networkcloud.models.VirtualMachineDeviceModelType;
import com.azure.resourcemanager.networkcloud.models.VirtualMachineIpAllocationMethod;
import com.azure.resourcemanager.networkcloud.models.VirtualMachinePlacementHint;
import com.azure.resourcemanager.networkcloud.models.VirtualMachinePlacementHintPodAffinityScope;
import com.azure.resourcemanager.networkcloud.models.VirtualMachinePlacementHintType;
import com.azure.resourcemanager.networkcloud.models.VirtualMachineSchedulingExecution;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualMachines CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/
     * VirtualMachines_Create.json
     */
    /**
     * Sample code: Create or update virtual machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        createOrUpdateVirtualMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.virtualMachines().define("virtualMachineName").withRegion("location")
            .withExistingResourceGroup("resourceGroupName")
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName")
                .withType("CustomLocation"))
            .withAdminUsername("username")
            .withCloudServicesNetworkAttachment(new NetworkAttachment().withAttachedNetworkId(
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/cloudServicesNetworks/cloudServicesNetworkName")
                .withIpAllocationMethod(VirtualMachineIpAllocationMethod.DYNAMIC))
            .withCpuCores(2L).withMemorySizeGB(8L)
            .withStorageProfile(new StorageProfile()
                .withOsDisk(new OsDisk().withCreateOption(OsDiskCreateOption.EPHEMERAL)
                    .withDeleteOption(OsDiskDeleteOption.DELETE).withDiskSizeGB(120L))
                .withVolumeAttachments(Arrays.asList(
                    "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/volumes/volumeName")))
            .withVmImage("myacr.azurecr.io/foobar:latest")
            .withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withBootMethod(VirtualMachineBootMethod.UEFI)
            .withNetworkAttachments(Arrays.asList(new NetworkAttachment().withAttachedNetworkId(
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName")
                .withDefaultGateway(DefaultGateway.TRUE)
                .withIpAllocationMethod(VirtualMachineIpAllocationMethod.DYNAMIC).withIpv4Address("198.51.100.1")
                .withIpv6Address("2001:0db8:0000:0000:0000:0000:0000:0000")
                .withNetworkAttachmentName("netAttachName01")))
            .withNetworkData("bmV0d29ya0RhdGVTYW1wbGU=")
            .withPlacementHints(Arrays.asList(new VirtualMachinePlacementHint()
                .withHintType(VirtualMachinePlacementHintType.AFFINITY)
                .withResourceId(
                    "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName")
                .withSchedulingExecution(VirtualMachineSchedulingExecution.HARD)
                .withScope(VirtualMachinePlacementHintPodAffinityScope.fromString(""))))
            .withSshPublicKeys(Arrays.asList(new SshPublicKey().withKeyData("fakeTokenPlaceholder")))
            .withUserData("dXNlckRhdGVTYW1wbGU=").withVmDeviceModel(VirtualMachineDeviceModelType.T2)
            .withVmImageRepositoryCredentials(new ImageRepositoryCredentials().withPassword("fakeTokenPlaceholder")
                .withRegistryUrl("myacr.azurecr.io").withUsername("myuser"))
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
