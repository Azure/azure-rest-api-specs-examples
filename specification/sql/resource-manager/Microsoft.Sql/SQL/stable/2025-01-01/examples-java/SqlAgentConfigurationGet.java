
/**
 * Samples for SqlAgent Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SqlAgentConfigurationGet.json
     */
    /**
     * Sample code: Gets current instance sql agent configuration.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsCurrentInstanceSqlAgentConfiguration(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSqlAgents().getWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645",
            com.azure.core.util.Context.NONE);
    }
}
