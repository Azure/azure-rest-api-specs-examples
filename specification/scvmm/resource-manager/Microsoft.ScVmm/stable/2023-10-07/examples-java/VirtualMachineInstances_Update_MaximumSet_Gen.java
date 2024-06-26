
import com.azure.resourcemanager.scvmm.models.AllocationMethod;
import com.azure.resourcemanager.scvmm.models.AvailabilitySetListItem;
import com.azure.resourcemanager.scvmm.models.DynamicMemoryEnabled;
import com.azure.resourcemanager.scvmm.models.HardwareProfileUpdate;
import com.azure.resourcemanager.scvmm.models.InfrastructureProfileUpdate;
import com.azure.resourcemanager.scvmm.models.LimitCpuForMigration;
import com.azure.resourcemanager.scvmm.models.NetworkInterfaceUpdate;
import com.azure.resourcemanager.scvmm.models.NetworkProfileUpdate;
import com.azure.resourcemanager.scvmm.models.StorageProfileUpdate;
import com.azure.resourcemanager.scvmm.models.StorageQosPolicyDetails;
import com.azure.resourcemanager.scvmm.models.VirtualDiskUpdate;
import com.azure.resourcemanager.scvmm.models.VirtualMachineInstanceUpdate;
import com.azure.resourcemanager.scvmm.models.VirtualMachineInstanceUpdateProperties;
import java.util.Arrays;

/**
 * Samples for VirtualMachineInstances Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualMachineInstances_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineInstances_Update_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualMachineInstancesUpdateMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineInstances().update("gtgclehcbsyave",
            new VirtualMachineInstanceUpdate().withProperties(new VirtualMachineInstanceUpdateProperties()
                .withAvailabilitySets(Arrays.asList(new AvailabilitySetListItem().withId(
                    "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/availabilitySets/availabilitySetResourceName")
                    .withName("lwbhaseo")))
                .withHardwareProfile(new HardwareProfileUpdate().withMemoryMB(5).withCpuCount(22)
                    .withLimitCpuForMigration(LimitCpuForMigration.TRUE)
                    .withDynamicMemoryEnabled(DynamicMemoryEnabled.TRUE).withDynamicMemoryMaxMB(2)
                    .withDynamicMemoryMinMB(30))
                .withNetworkProfile(new NetworkProfileUpdate().withNetworkInterfaces(
                    Arrays.asList(new NetworkInterfaceUpdate().withName("kvofzqulbjlbtt").withMacAddress("oaeqqegt")
                        .withVirtualNetworkId(
                            "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualNetworks/virtualNetworkName")
                        .withIpv4AddressType(AllocationMethod.DYNAMIC).withIpv6AddressType(AllocationMethod.DYNAMIC)
                        .withMacAddressType(AllocationMethod.DYNAMIC).withNicId("roxpsvlo"))))
                .withStorageProfile(new StorageProfileUpdate().withDisks(Arrays.asList(new VirtualDiskUpdate()
                    .withName("fgnckfymwdsqnfxkdvexuaobe").withDiskId("ltdrwcfjklpsimhzqyh").withDiskSizeGB(30)
                    .withBus(8).withLun(10).withBusType("zu").withVhdType("cnbeeeylrvopigdynvgpkfp")
                    .withStorageQosPolicy(new StorageQosPolicyDetails().withName("ceiyfrflu").withId("o")))))
                .withInfrastructureProfile(
                    new InfrastructureProfileUpdate().withCheckpointType("jkbpzjxpeegackhsvikrnlnwqz"))),
            com.azure.core.util.Context.NONE);
    }
}
