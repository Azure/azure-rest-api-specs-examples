
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ElasticPoolUpdate;

/** Samples for ElasticPools Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ElasticPoolUpdateResetMaintenanceConfiguration.json
     */
    /**
     * Sample code: Resets maintenance configuration of an elastic pool to default.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        resetsMaintenanceConfigurationOfAnElasticPoolToDefault(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getElasticPools().update("sqlcrudtest-2369", "sqlcrudtest-8069",
            "sqlcrudtest-8102",
            new ElasticPoolUpdate().withMaintenanceConfigurationId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_Default"),
            Context.NONE);
    }
}
