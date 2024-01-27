
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.EncryptionProtectorInner;
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;
import com.azure.resourcemanager.sql.models.ServerKeyType;

/** Samples for EncryptionProtectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * EncryptionProtectorCreateOrUpdateKeyVault.json
     */
    /**
     * Sample code: Update the encryption protector to key vault.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTheEncryptionProtectorToKeyVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getEncryptionProtectors().createOrUpdate("sqlcrudtest-7398",
            "sqlcrudtest-4645", EncryptionProtectorName.CURRENT,
            new EncryptionProtectorInner().withServerKeyName("fakeTokenPlaceholder")
                .withServerKeyType(ServerKeyType.AZURE_KEY_VAULT).withAutoRotationEnabled(false),
            Context.NONE);
    }
}
