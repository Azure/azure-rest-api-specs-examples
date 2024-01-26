
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ElasticPoolInner;

/** Samples for ElasticPools CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ElasticPoolCreateOrUpdateSetMaintenanceConfiguration.json
     */
    /**
     * Sample code: Create or update elastic pool with maintenance configuration parameter.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateElasticPoolWithMaintenanceConfigurationParameter(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getElasticPools().createOrUpdate("sqlcrudtest-2369",
            "sqlcrudtest-8069", "sqlcrudtest-8102",
            new ElasticPoolInner().withLocation("Japan East").withMaintenanceConfigurationId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_1"),
            Context.NONE);
    }
}
