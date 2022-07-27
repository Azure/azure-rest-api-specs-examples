import com.azure.core.util.Context;

/** Samples for ManagedInstanceKeys ListByInstance. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/ManagedInstanceKeyList.json
     */
    /**
     * Sample code: List the keys for a managed instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTheKeysForAManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedInstanceKeys()
            .listByInstance("sqlcrudtest-7398", "sqlcrudtest-4645", null, Context.NONE);
    }
}
