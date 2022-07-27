import com.azure.core.util.Context;

/** Samples for ManagedInstances GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/ManagedInstanceGet.json
     */
    /**
     * Sample code: Get managed instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedInstances()
            .getByResourceGroupWithResponse("testrg", "testinstance", Context.NONE);
    }
}
