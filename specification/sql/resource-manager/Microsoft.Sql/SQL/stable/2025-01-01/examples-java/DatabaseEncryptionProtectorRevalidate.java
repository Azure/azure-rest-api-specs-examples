
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;

/**
 * Samples for DatabaseEncryptionProtectors Revalidate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseEncryptionProtectorRevalidate.json
     */
    /**
     * Sample code: Revalidates the encryption protector for a particular database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void revalidatesTheEncryptionProtectorForAParticularDatabase(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseEncryptionProtectors().revalidate("sqlcrudtest-7398", "sqlcrudtest-4645",
            "sqlcrudtestdb-2342", EncryptionProtectorName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
