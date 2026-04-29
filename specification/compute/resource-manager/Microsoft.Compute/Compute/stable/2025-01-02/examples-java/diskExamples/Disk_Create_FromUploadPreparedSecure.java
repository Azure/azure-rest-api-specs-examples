
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.DiskSecurityProfile;
import com.azure.resourcemanager.compute.models.DiskSecurityTypes;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskExamples/Disk_Create_FromUploadPreparedSecure.json
     */
    /**
     * Sample code: create a managed disk from UploadPreparedSecure create option.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createAManagedDiskFromUploadPreparedSecureCreateOption(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk",
            new DiskInner().withLocation("West US").withOsType(OperatingSystemTypes.WINDOWS)
                .withCreationData(new CreationData().withCreateOption(DiskCreateOption.UPLOAD_PREPARED_SECURE)
                    .withUploadSizeBytes(10737418752L))
                .withSecurityProfile(new DiskSecurityProfile().withSecurityType(DiskSecurityTypes.TRUSTED_LAUNCH)),
            com.azure.core.util.Context.NONE);
    }
}
