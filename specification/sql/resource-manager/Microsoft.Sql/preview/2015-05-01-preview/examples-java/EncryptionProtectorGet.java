import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;

/** Samples for EncryptionProtectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2015-05-01-preview/examples/EncryptionProtectorGet.json
     */
    /**
     * Sample code: Get the encryption protector.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheEncryptionProtector(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getEncryptionProtectors()
            .getWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645", EncryptionProtectorName.CURRENT, Context.NONE);
    }
}
