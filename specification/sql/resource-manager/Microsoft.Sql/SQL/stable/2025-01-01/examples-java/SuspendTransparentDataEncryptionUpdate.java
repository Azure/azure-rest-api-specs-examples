
import com.azure.resourcemanager.sql.models.TransparentDataEncryptionName;

/**
 * Samples for TransparentDataEncryptions Suspend.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SuspendTransparentDataEncryptionUpdate.json
     */
    /**
     * Sample code: Suspend database's Transparent Data Encryption scan state with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void suspendDatabaseSTransparentDataEncryptionScanStateWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getTransparentDataEncryptions().suspend("securitytde-42-rg", "securitytde-42", "testdb",
            TransparentDataEncryptionName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
