
import com.azure.resourcemanager.sql.models.TransparentDataEncryptionName;

/**
 * Samples for ManagedDatabaseTransparentDataEncryption Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedTransparentDataEncryptionGet.json
     */
    /**
     * Sample code: Get a database's transparent data encryption.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getADatabaseSTransparentDataEncryption(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseTransparentDataEncryptions().getWithResponse(
            "security-tde-resourcegroup", "securitytde", "testdb", TransparentDataEncryptionName.CURRENT,
            com.azure.core.util.Context.NONE);
    }
}
