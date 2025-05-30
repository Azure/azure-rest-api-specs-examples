
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineUpdateInner;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DataDisk;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.DiskDetachOptionTypes;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.ManagedDiskParameters;
import com.azure.resourcemanager.compute.models.NetworkInterfaceReference;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSProfile;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.StorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineSizeTypes;
import java.util.Arrays;

/**
 * Samples for VirtualMachines Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineExamples/VirtualMachine_Update_ForceDetachDataDisk.json
     */
    /**
     * Sample code: Update a VM by force-detaching data disk.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAVMByForceDetachingDataDisk(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().update("myResourceGroup", "myVM",
            new VirtualMachineUpdateInner()
                .withHardwareProfile(new HardwareProfile().withVmSize(VirtualMachineSizeTypes.STANDARD_D2_V2))
                .withStorageProfile(new StorageProfile()
                    .withImageReference(new ImageReference().withPublisher("MicrosoftWindowsServer")
                        .withOffer("WindowsServer").withSku("2016-Datacenter").withVersion("latest"))
                    .withOsDisk(new OSDisk().withName("myVMosdisk").withCaching(CachingTypes.READ_WRITE)
                        .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE).withManagedDisk(
                            new ManagedDiskParameters().withStorageAccountType(StorageAccountTypes.STANDARD_LRS)))
                    .withDataDisks(Arrays.asList(
                        new DataDisk().withLun(0).withCreateOption(DiskCreateOptionTypes.EMPTY).withDiskSizeGB(1023)
                            .withToBeDetached(true).withDetachOption(DiskDetachOptionTypes.FORCE_DETACH),
                        new DataDisk().withLun(1).withCreateOption(DiskCreateOptionTypes.EMPTY).withDiskSizeGB(1023)
                            .withToBeDetached(false))))
                .withOsProfile(new OSProfile().withComputerName("myVM").withAdminUsername("{your-username}")
                    .withAdminPassword("fakeTokenPlaceholder"))
                .withNetworkProfile(
                    new NetworkProfile().withNetworkInterfaces(Arrays.asList(new NetworkInterfaceReference().withId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}")
                        .withPrimary(true)))),
            null, null, com.azure.core.util.Context.NONE);
    }
}
