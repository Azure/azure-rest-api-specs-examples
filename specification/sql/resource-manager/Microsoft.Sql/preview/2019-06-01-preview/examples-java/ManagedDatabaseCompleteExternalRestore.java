import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.CompleteDatabaseRestoreDefinition;

/** Samples for ManagedDatabases CompleteRestore. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/ManagedDatabaseCompleteExternalRestore.json
     */
    /**
     * Sample code: Completes a managed database external backup restore.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void completesAManagedDatabaseExternalBackupRestore(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedDatabases()
            .completeRestore(
                "myRG",
                "myManagedInstanceName",
                "myDatabase",
                new CompleteDatabaseRestoreDefinition().withLastBackupName("testdb1_log4"),
                Context.NONE);
    }
}
