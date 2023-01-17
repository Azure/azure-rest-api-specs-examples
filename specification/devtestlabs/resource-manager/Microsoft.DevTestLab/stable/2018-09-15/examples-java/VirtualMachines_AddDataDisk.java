import com.azure.resourcemanager.devtestlabs.models.AttachNewDataDiskOptions;
import com.azure.resourcemanager.devtestlabs.models.DataDiskProperties;
import com.azure.resourcemanager.devtestlabs.models.StorageType;

/** Samples for VirtualMachines AddDataDisk. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_AddDataDisk.json
     */
    /**
     * Sample code: VirtualMachines_AddDataDisk.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesAddDataDisk(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .virtualMachines()
            .addDataDisk(
                "resourceGroupName",
                "{labName}",
                "{virtualMachineName}",
                new DataDiskProperties()
                    .withAttachNewDataDiskOptions(
                        new AttachNewDataDiskOptions()
                            .withDiskSizeGiB(127)
                            .withDiskName("{diskName}")
                            .withDiskType(StorageType.fromString("{diskType}"))),
                com.azure.core.util.Context.NONE);
    }
}
