
import com.azure.resourcemanager.sql.models.ElasticPoolPerDatabaseSettings;
import com.azure.resourcemanager.sql.models.ElasticPoolUpdate;
import com.azure.resourcemanager.sql.models.Sku;

/**
 * Samples for ElasticPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ElasticPoolUpdateServerlessProperties.json
     */
    /**
     * Sample code: Update an elastic pool with serverless properties.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        updateAnElasticPoolWithServerlessProperties(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().update("sqlcrudtest-2369", "sqlcrudtest-8069", "sqlcrudtest-8102",
            new ElasticPoolUpdate()
                .withSku(new Sku().withName("GP_S_Gen5_2").withTier("GeneralPurpose").withCapacity(2))
                .withMinCapacity(0.5D)
                .withPerDatabaseSettings(new ElasticPoolPerDatabaseSettings().withMinCapacity(0.0D)
                    .withMaxCapacity(2.0D).withAutoPauseDelay(80))
                .withAutoPauseDelay(60),
            com.azure.core.util.Context.NONE);
    }
}
