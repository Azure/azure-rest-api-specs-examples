import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ElasticPoolLicenseType;
import com.azure.resourcemanager.sql.models.ElasticPoolPerDatabaseSettings;
import com.azure.resourcemanager.sql.models.ElasticPoolUpdate;
import com.azure.resourcemanager.sql.models.Sku;

/** Samples for ElasticPools Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/ElasticPoolUpdateMax.json
     */
    /**
     * Sample code: Update an elastic pool with all parameter.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAnElasticPoolWithAllParameter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getElasticPools()
            .update(
                "sqlcrudtest-2369",
                "sqlcrudtest-8069",
                "sqlcrudtest-8102",
                new ElasticPoolUpdate()
                    .withSku(new Sku().withName("BC_Gen4_2").withTier("BusinessCritical").withCapacity(2))
                    .withPerDatabaseSettings(
                        new ElasticPoolPerDatabaseSettings().withMinCapacity(0.25).withMaxCapacity(1.0))
                    .withZoneRedundant(true)
                    .withLicenseType(ElasticPoolLicenseType.LICENSE_INCLUDED),
                Context.NONE);
    }
}
