
import com.azure.resourcemanager.sql.models.TransparentDataEncryptionName;

/**
 * Samples for TransparentDataEncryptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/TransparentDataEncryptionGet.json
     */
    /**
     * Sample code: Get a database's transparent data encryption.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getADatabaseSTransparentDataEncryption(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getTransparentDataEncryptions().getWithResponse(
            "security-tde-resourcegroup", "securitytde", "testdb", TransparentDataEncryptionName.CURRENT,
            com.azure.core.util.Context.NONE);
    }
}
