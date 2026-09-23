
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;

/**
 * Samples for ManagedInstanceEncryptionProtectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedInstanceEncryptionProtectorGet.json
     */
    /**
     * Sample code: Get the encryption protector.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getTheEncryptionProtector(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceEncryptionProtectors().getWithResponse("sqlcrudtest-7398",
            "sqlcrudtest-4645", EncryptionProtectorName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
