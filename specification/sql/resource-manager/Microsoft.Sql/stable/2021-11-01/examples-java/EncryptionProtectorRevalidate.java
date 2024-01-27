
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;

/** Samples for EncryptionProtectors Revalidate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/EncryptionProtectorRevalidate.json
     */
    /**
     * Sample code: Revalidates the encryption protector.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void revalidatesTheEncryptionProtector(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getEncryptionProtectors().revalidate("sqlcrudtest-7398",
            "sqlcrudtest-4645", EncryptionProtectorName.CURRENT, Context.NONE);
    }
}
