
import com.azure.resourcemanager.recoveryservicesbackup.models.AzureFileShareRestoreRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.CopyOptions;
import com.azure.resourcemanager.recoveryservicesbackup.models.IdentityInfo;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RestoreRequestResource;
import com.azure.resourcemanager.recoveryservicesbackup.models.RestoreRequestType;

/**
 * Samples for Restores Trigger.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/AzureStorage/TriggerRestore_AzureFileShare_WithSAMI.json
     */
    /**
     * Sample code: Restore Azure File Share to Original Location with Managed Identity.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void restoreAzureFileShareToOriginalLocationWithManagedIdentity(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.restores().trigger("swaggertestvault", "SwaggerTestRg", "Azure",
            "StorageContainer;Storage;SwaggerTestRg;swaggertestsa", "AzureFileShare;testshare", "932886657837421071",
            new RestoreRequestResource().withProperties(new AzureFileShareRestoreRequest()
                .withRecoveryType(RecoveryType.ORIGINAL_LOCATION)
                .withSourceResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsa")
                .withCopyOptions(CopyOptions.OVERWRITE).withRestoreRequestType(RestoreRequestType.FULL_SHARE_RESTORE)
                .withIdentityInfo(new IdentityInfo().withIsSystemAssignedIdentity(true))),
            com.azure.core.util.Context.NONE);
    }
}
