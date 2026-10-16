
import com.azure.resourcemanager.azurestackhci.models.DeleteVolumeJobProperties;
import com.azure.resourcemanager.azurestackhci.models.DeploymentMode;

/**
 * Samples for EdgeMachineDiskPrivilegedJobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineDiskPrivilegedJobs_CreateOrUpdate_DeleteVolume.json
     */
    /**
     * Sample code: EdgeMachineDiskPrivilegedJobs_CreateOrUpdate_DeleteVolume.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeMachineDiskPrivilegedJobsCreateOrUpdateDeleteVolume(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineDiskPrivilegedJobs().define("DeleteVolume")
            .withExistingDisk("test-rg", "EdgeMachine01", "disk-001").withProperties(new DeleteVolumeJobProperties()
                .withDeploymentMode(DeploymentMode.DEPLOY).withVolumePath("/data").withConfirmDeletion(true))
            .create();
    }
}
