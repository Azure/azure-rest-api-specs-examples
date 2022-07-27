import com.azure.core.util.Context;

/** Samples for Usages ListByInstancePool. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/ListInstancePoolUsage.json
     */
    /**
     * Sample code: List instance pool usages.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listInstancePoolUsages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getUsages()
            .listByInstancePool("group1", "testIP", null, Context.NONE);
    }
}
