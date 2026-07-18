
import com.azure.resourcemanager.sql.models.CompleteDatabaseRestoreDefinition;

/**
 * Samples for ManagedDatabases CompleteRestore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseCompleteExternalRestore.json
     */
    /**
     * Sample code: Completes a managed database external backup restore.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        completesAManagedDatabaseExternalBackupRestore(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabases().completeRestore("myRG", "myManagedInstanceName", "myDatabase",
            new CompleteDatabaseRestoreDefinition().withLastBackupName("testdb1_log4"),
            com.azure.core.util.Context.NONE);
    }
}
