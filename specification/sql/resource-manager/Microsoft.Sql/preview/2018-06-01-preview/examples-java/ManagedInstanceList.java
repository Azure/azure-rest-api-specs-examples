import com.azure.core.util.Context;

/** Samples for ManagedInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/ManagedInstanceList.json
     */
    /**
     * Sample code: List managed instances.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedInstances(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().list(Context.NONE);
    }
}
