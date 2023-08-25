import com.azure.resourcemanager.apimanagement.models.AccessType;
import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceBackupRestoreParameters;

/** Samples for ApiManagementService Backup. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementBackupWithSystemManagedIdentity.json
     */
    /**
     * Sample code: ApiManagementBackupWithSystemManagedIdentity.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementBackupWithSystemManagedIdentity(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServices()
            .backup(
                "rg1",
                "apimService1",
                new ApiManagementServiceBackupRestoreParameters()
                    .withStorageAccount("contosorpstorage")
                    .withContainerName("apim-backups")
                    .withBackupName("backup5")
                    .withAccessType(AccessType.SYSTEM_ASSIGNED_MANAGED_IDENTITY),
                com.azure.core.util.Context.NONE);
    }
}
