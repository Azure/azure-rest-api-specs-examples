
import com.azure.resourcemanager.apimanagement.models.AccessType;
import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceBackupRestoreParameters;

/**
 * Samples for ApiManagementService Restore.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementRestoreWithAccessKey.json
     */
    /**
     * Sample code: ApiManagementRestoreService.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementRestoreService(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementServices().restore("rg1", "apimService1",
            new ApiManagementServiceBackupRestoreParameters().withStorageAccount("teststorageaccount")
                .withContainerName("backupContainer").withBackupName("apimService1backup_2017_03_19")
                .withAccessType(AccessType.ACCESS_KEY).withAccessKey("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
