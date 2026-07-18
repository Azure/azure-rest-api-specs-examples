
import com.azure.resourcemanager.sql.fluent.models.SqlAgentConfigurationInner;
import com.azure.resourcemanager.sql.models.SqlAgentConfigurationPropertiesState;

/**
 * Samples for SqlAgent CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SqlAgentConfigurationPut.json
     */
    /**
     * Sample code: Puts new sql agent configuration to instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void putsNewSqlAgentConfigurationToInstance(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSqlAgents().createOrUpdateWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645",
            new SqlAgentConfigurationInner().withState(SqlAgentConfigurationPropertiesState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
