
import com.azure.resourcemanager.compute.models.AccessLevel;
import com.azure.resourcemanager.compute.models.FileFormat;
import com.azure.resourcemanager.compute.models.GrantAccessData;

/**
 * Samples for DiskRestorePoint GrantAccess.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskRestorePointExamples/DiskRestorePoint_BeginGetAccess.json
     */
    /**
     * Sample code: grants access to a diskRestorePoint.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void grantsAccessToADiskRestorePoint(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskRestorePoints().grantAccess("myResourceGroup", "rpc", "vmrp",
            "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745", new GrantAccessData()
                .withAccess(AccessLevel.READ).withDurationInSeconds(300).withFileFormat(FileFormat.VHDX),
            com.azure.core.util.Context.NONE);
    }
}
