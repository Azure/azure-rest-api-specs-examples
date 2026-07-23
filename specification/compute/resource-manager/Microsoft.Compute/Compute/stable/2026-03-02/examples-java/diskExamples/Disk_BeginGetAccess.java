
import com.azure.resourcemanager.compute.models.AccessLevel;
import com.azure.resourcemanager.compute.models.FileFormat;
import com.azure.resourcemanager.compute.models.GrantAccessData;

/**
 * Samples for Disks GrantAccess.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_BeginGetAccess.json
     */
    /**
     * Sample code: get a sas on a managed disk.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getASasOnAManagedDisk(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().grantAccess("myResourceGroup", "myDisk", new GrantAccessData()
            .withAccess(AccessLevel.READ).withDurationInSeconds(300).withFileFormat(FileFormat.VHD),
            com.azure.core.util.Context.NONE);
    }
}
