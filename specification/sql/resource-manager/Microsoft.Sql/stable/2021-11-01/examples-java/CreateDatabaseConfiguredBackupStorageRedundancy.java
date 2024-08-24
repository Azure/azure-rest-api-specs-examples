
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.BackupStorageRedundancy;

/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * CreateDatabaseConfiguredBackupStorageRedundancy.json
     */
    /**
     * Sample code: Creates a database with specified backup storage redundancy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createsADatabaseWithSpecifiedBackupStorageRedundancy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers().manager().serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia",
                "testsvr", "testdb", new DatabaseInner().withLocation("southeastasia")
                    .withRequestedBackupStorageRedundancy(BackupStorageRedundancy.ZONE),
                com.azure.core.util.Context.NONE);
    }
}
