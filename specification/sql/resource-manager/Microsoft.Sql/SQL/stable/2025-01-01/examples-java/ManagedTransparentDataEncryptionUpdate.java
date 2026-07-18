
import com.azure.resourcemanager.sql.fluent.models.ManagedTransparentDataEncryptionInner;
import com.azure.resourcemanager.sql.models.TransparentDataEncryptionName;
import com.azure.resourcemanager.sql.models.TransparentDataEncryptionState;

/**
 * Samples for ManagedDatabaseTransparentDataEncryption CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedTransparentDataEncryptionUpdate.json
     */
    /**
     * Sample code: Update a database's Transparent Data Encryption state with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateADatabaseSTransparentDataEncryptionStateWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseTransparentDataEncryptions().createOrUpdateWithResponse(
            "securitytde-42-rg", "securitytde-42", "testdb", TransparentDataEncryptionName.CURRENT,
            new ManagedTransparentDataEncryptionInner().withState(TransparentDataEncryptionState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
