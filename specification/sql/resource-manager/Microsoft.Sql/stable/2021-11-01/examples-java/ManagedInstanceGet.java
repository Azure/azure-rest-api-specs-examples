
/**
 * Samples for ManagedInstances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceGet.json
     */
    /**
     * Sample code: Get managed instance.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().getByResourceGroupWithResponse("testrg",
            "testinstance", null, com.azure.core.util.Context.NONE);
    }
}
