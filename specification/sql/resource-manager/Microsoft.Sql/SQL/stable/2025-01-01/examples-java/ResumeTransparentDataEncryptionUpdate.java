
import com.azure.resourcemanager.sql.models.TransparentDataEncryptionName;

/**
 * Samples for TransparentDataEncryptions Resume.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ResumeTransparentDataEncryptionUpdate.json
     */
    /**
     * Sample code: Resume database's Transparent Data Encryption scan state with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void resumeDatabaseSTransparentDataEncryptionScanStateWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getTransparentDataEncryptions().resume("securitytde-42-rg", "securitytde-42", "testdb",
            TransparentDataEncryptionName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
