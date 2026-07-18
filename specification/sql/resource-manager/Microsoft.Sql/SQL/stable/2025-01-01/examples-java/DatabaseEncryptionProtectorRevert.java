
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;

/**
 * Samples for DatabaseEncryptionProtectors Revert.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseEncryptionProtectorRevert.json
     */
    /**
     * Sample code: Reverts the encryption protector for a particular database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        revertsTheEncryptionProtectorForAParticularDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseEncryptionProtectors().revert("sqlcrudtest-7398", "sqlcrudtest-4645",
            "sqlcrudtestdb-2342", EncryptionProtectorName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
