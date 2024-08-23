
/**
 * Samples for SqlAgent Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SqlAgentConfigurationGet.json
     */
    /**
     * Sample code: Gets current instance sql agent configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsCurrentInstanceSqlAgentConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSqlAgents().getWithResponse("sqlcrudtest-7398",
            "sqlcrudtest-4645", com.azure.core.util.Context.NONE);
    }
}
