import com.azure.core.util.Context;

/** Samples for ManagedInstances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/ManagedInstanceDelete.json
     */
    /**
     * Sample code: Delete managed instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedInstances()
            .delete("testrg", "testinstance", Context.NONE);
    }
}
