
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;

/**
 * Samples for EncryptionProtectors Revalidate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/EncryptionProtectorRevalidate.json
     */
    /**
     * Sample code: Revalidates the encryption protector.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void revalidatesTheEncryptionProtector(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getEncryptionProtectors().revalidate("sqlcrudtest-7398", "sqlcrudtest-4645",
            EncryptionProtectorName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
