
import com.azure.resourcemanager.sql.fluent.models.LogicalDatabaseTransparentDataEncryptionInner;
import com.azure.resourcemanager.sql.models.TransparentDataEncryptionName;
import com.azure.resourcemanager.sql.models.TransparentDataEncryptionScanState;
import com.azure.resourcemanager.sql.models.TransparentDataEncryptionState;

/**
 * Samples for TransparentDataEncryptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/TransparentDataEncryptionUpdate.json
     */
    /**
     * Sample code: Update a database's Transparent Data Encryption state with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateADatabaseSTransparentDataEncryptionStateWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getTransparentDataEncryptions().createOrUpdate("securitytde-42-rg", "securitytde-42",
            "testdb", TransparentDataEncryptionName.CURRENT,
            new LogicalDatabaseTransparentDataEncryptionInner().withState(TransparentDataEncryptionState.ENABLED)
                .withScanState(TransparentDataEncryptionScanState.SUSPEND),
            com.azure.core.util.Context.NONE);
    }
}
