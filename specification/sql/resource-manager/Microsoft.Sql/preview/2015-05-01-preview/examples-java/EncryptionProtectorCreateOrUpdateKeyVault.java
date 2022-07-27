import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.EncryptionProtectorInner;
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;
import com.azure.resourcemanager.sql.models.ServerKeyType;

/** Samples for EncryptionProtectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2015-05-01-preview/examples/EncryptionProtectorCreateOrUpdateKeyVault.json
     */
    /**
     * Sample code: Update the encryption protector to key vault.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTheEncryptionProtectorToKeyVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getEncryptionProtectors()
            .createOrUpdate(
                "sqlcrudtest-7398",
                "sqlcrudtest-4645",
                EncryptionProtectorName.CURRENT,
                new EncryptionProtectorInner()
                    .withServerKeyName("someVault_someKey_01234567890123456789012345678901")
                    .withServerKeyType(ServerKeyType.AZURE_KEY_VAULT),
                Context.NONE);
    }
}
