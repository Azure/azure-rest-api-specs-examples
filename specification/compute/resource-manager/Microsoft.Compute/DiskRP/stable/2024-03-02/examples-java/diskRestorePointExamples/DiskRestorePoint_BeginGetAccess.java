
import com.azure.resourcemanager.compute.models.AccessLevel;
import com.azure.resourcemanager.compute.models.FileFormat;
import com.azure.resourcemanager.compute.models.GrantAccessData;

/**
 * Samples for DiskRestorePoint GrantAccess.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/
     * diskRestorePointExamples/DiskRestorePoint_BeginGetAccess.json
     */
    /**
     * Sample code: Grants access to a diskRestorePoint.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void grantsAccessToADiskRestorePoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDiskRestorePoints().grantAccess("myResourceGroup", "rpc",
            "vmrp", "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745", new GrantAccessData()
                .withAccess(AccessLevel.READ).withDurationInSeconds(300).withFileFormat(FileFormat.VHDX),
            com.azure.core.util.Context.NONE);
    }
}
