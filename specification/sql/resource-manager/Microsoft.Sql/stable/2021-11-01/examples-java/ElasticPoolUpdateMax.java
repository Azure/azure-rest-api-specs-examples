
import com.azure.resourcemanager.sql.models.ElasticPoolLicenseType;
import com.azure.resourcemanager.sql.models.ElasticPoolPerDatabaseSettings;
import com.azure.resourcemanager.sql.models.ElasticPoolUpdate;
import com.azure.resourcemanager.sql.models.Sku;

/**
 * Samples for ElasticPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ElasticPoolUpdateMax.json
     */
    /**
     * Sample code: Update an elastic pool with all parameter.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAnElasticPoolWithAllParameter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getElasticPools().update("sqlcrudtest-2369", "sqlcrudtest-8069",
            "sqlcrudtest-8102",
            new ElasticPoolUpdate().withSku(new Sku().withName("BC_Gen4").withTier("BusinessCritical").withCapacity(2))
                .withPerDatabaseSettings(
                    new ElasticPoolPerDatabaseSettings().withMinCapacity(0.25D).withMaxCapacity(1.0D))
                .withZoneRedundant(true).withLicenseType(ElasticPoolLicenseType.LICENSE_INCLUDED),
            com.azure.core.util.Context.NONE);
    }
}
