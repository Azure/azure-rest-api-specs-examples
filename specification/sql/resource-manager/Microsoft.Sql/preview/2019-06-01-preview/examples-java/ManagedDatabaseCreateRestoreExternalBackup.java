import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ManagedDatabaseInner;
import com.azure.resourcemanager.sql.models.ManagedDatabaseCreateMode;
import java.util.HashMap;
import java.util.Map;

/** Samples for ManagedDatabases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/ManagedDatabaseCreateRestoreExternalBackup.json
     */
    /**
     * Sample code: Creates a new managed database by restoring from an external backup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsANewManagedDatabaseByRestoringFromAnExternalBackup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedDatabases()
            .createOrUpdate(
                "Default-SQL-SouthEastAsia",
                "managedInstance",
                "managedDatabase",
                new ManagedDatabaseInner()
                    .withLocation("southeastasia")
                    .withCollation("SQL_Latin1_General_CP1_CI_AS")
                    .withCreateMode(ManagedDatabaseCreateMode.RESTORE_EXTERNAL_BACKUP)
                    .withStorageContainerUri("https://myaccountname.blob.core.windows.net/backups")
                    .withStorageContainerSasToken("sv=2015-12-11&sr=c&sp=rl&sig=1234"),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
