import com.azure.resourcemanager.apimanagement.models.AccessType;
import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceBackupRestoreParameters;

/** Samples for ApiManagementService Backup. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementBackupWithAccessKey.json
     */
    /**
     * Sample code: ApiManagementBackupWithAccessKey.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementBackupWithAccessKey(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServices()
            .backup(
                "rg1",
                "apimService1",
                new ApiManagementServiceBackupRestoreParameters()
                    .withStorageAccount("teststorageaccount")
                    .withContainerName("backupContainer")
                    .withBackupName("apimService1backup_2017_03_19")
                    .withAccessType(AccessType.ACCESS_KEY)
                    .withAccessKey("fakeTokenPlaceholder"),
                com.azure.core.util.Context.NONE);
    }
}
