
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;

/**
 * Samples for ManagedInstanceEncryptionProtectors Revalidate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedInstanceEncryptionProtectorRevalidate.json
     */
    /**
     * Sample code: Revalidates the encryption protector.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void revalidatesTheEncryptionProtector(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceEncryptionProtectors().revalidate("sqlcrudtest-7398",
            "sqlcrudtest-4645", EncryptionProtectorName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
