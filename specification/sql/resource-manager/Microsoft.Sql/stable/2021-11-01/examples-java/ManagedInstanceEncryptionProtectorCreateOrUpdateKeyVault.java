
import com.azure.resourcemanager.sql.fluent.models.ManagedInstanceEncryptionProtectorInner;
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;
import com.azure.resourcemanager.sql.models.ServerKeyType;

/**
 * Samples for ManagedInstanceEncryptionProtectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstanceEncryptionProtectorCreateOrUpdateKeyVault.json
     */
    /**
     * Sample code: Update the encryption protector to key vault.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTheEncryptionProtectorToKeyVault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstanceEncryptionProtectors().createOrUpdate(
            "sqlcrudtest-7398", "sqlcrudtest-4645", EncryptionProtectorName.CURRENT,
            new ManagedInstanceEncryptionProtectorInner().withServerKeyName("fakeTokenPlaceholder")
                .withServerKeyType(ServerKeyType.AZURE_KEY_VAULT).withAutoRotationEnabled(false),
            com.azure.core.util.Context.NONE);
    }
}
