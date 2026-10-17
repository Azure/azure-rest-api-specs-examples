
import com.azure.resourcemanager.azurestackhci.models.CreateVolumeJobProperties;
import com.azure.resourcemanager.azurestackhci.models.DeploymentMode;
import com.azure.resourcemanager.azurestackhci.models.DiskFileSystemType;

/**
 * Samples for EdgeMachineDiskJobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineDiskJobs_CreateOrUpdate_CreateVolume.json
     */
    /**
     * Sample code: EdgeMachineDiskJobs_CreateOrUpdate_CreateVolume.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeMachineDiskJobsCreateOrUpdateCreateVolume(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineDiskJobs().define("CreateVolume").withExistingDisk("test-rg", "EdgeMachine01", "disk-001")
            .withProperties(new CreateVolumeJobProperties().withDeploymentMode(DeploymentMode.DEPLOY)
                .withSizeInBytes("107374182400").withPath("/data").withFileSystem(DiskFileSystemType.EXT4))
            .create();
    }
}
