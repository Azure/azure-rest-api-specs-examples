
import com.azure.core.util.Context;

/** Samples for ManagedInstanceEncryptionProtectors ListByInstance. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstanceEncryptionProtectorList.json
     */
    /**
     * Sample code: List encryption protectors by managed instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listEncryptionProtectorsByManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstanceEncryptionProtectors()
            .listByInstance("sqlcrudtest-7398", "sqlcrudtest-4645", Context.NONE);
    }
}
