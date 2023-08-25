import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/** Samples for Disks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/diskExamples/Disk_Create_UploadDisk.json
     */
    /**
     * Sample code: Create a managed upload disk.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAManagedUploadDisk(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDisks()
            .createOrUpdate(
                "myResourceGroup",
                "myDisk",
                new DiskInner()
                    .withLocation("West US")
                    .withCreationData(
                        new CreationData().withCreateOption(DiskCreateOption.UPLOAD).withUploadSizeBytes(10737418752L)),
                com.azure.core.util.Context.NONE);
    }
}
