
import com.azure.core.util.Context;

/** Samples for InstancePools Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DeleteInstancePool.json
     */
    /**
     * Sample code: Delete an instance pool.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAnInstancePool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getInstancePools().delete("group1", "testIP", Context.NONE);
    }
}
