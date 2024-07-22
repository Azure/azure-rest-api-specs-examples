
import com.azure.resourcemanager.dataprotection.models.SuspendBackupRequest;
import java.util.Arrays;

/**
 * Samples for BackupInstances SuspendBackups.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * BackupInstanceOperations/SuspendBackup_ResourceGuardEnabled.json
     */
    /**
     * Sample code: SuspendBackups with MUA.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void suspendBackupsWithMUA(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().suspendBackups("testrg", "testvault", "testbi",
            new SuspendBackupRequest().withResourceGuardOperationRequests(Arrays.asList(
                "/subscriptions/754ec39f-8d2a-44cf-bfbf-13107ac85c36/resourcegroups/mua-testing/providers/Microsoft.DataProtection/resourceGuards/gvjreddy-test-ecy-rg-reader/dppDisableSuspendBackupsRequests/default")),
            com.azure.core.util.Context.NONE);
    }
}
