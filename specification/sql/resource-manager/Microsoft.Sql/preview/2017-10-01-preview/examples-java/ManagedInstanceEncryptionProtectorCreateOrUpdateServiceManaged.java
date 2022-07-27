import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ManagedInstanceEncryptionProtectorInner;
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;
import com.azure.resourcemanager.sql.models.ServerKeyType;

/** Samples for ManagedInstanceEncryptionProtectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/ManagedInstanceEncryptionProtectorCreateOrUpdateServiceManaged.json
     */
    /**
     * Sample code: Update the encryption protector to service managed.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTheEncryptionProtectorToServiceManaged(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedInstanceEncryptionProtectors()
            .createOrUpdate(
                "sqlcrudtest-7398",
                "sqlcrudtest-4645",
                EncryptionProtectorName.CURRENT,
                new ManagedInstanceEncryptionProtectorInner()
                    .withServerKeyName("ServiceManaged")
                    .withServerKeyType(ServerKeyType.SERVICE_MANAGED),
                Context.NONE);
    }
}
