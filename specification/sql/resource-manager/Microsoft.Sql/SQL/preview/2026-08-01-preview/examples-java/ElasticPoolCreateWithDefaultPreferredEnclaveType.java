
import com.azure.resourcemanager.sql.fluent.models.ElasticPoolInner;
import com.azure.resourcemanager.sql.models.AlwaysEncryptedEnclaveType;
import com.azure.resourcemanager.sql.models.Sku;

/**
 * Samples for ElasticPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ElasticPoolCreateWithDefaultPreferredEnclaveType.json
     */
    /**
     * Sample code: Create or update elastic pool with preferred enclave type parameter as Default.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createOrUpdateElasticPoolWithPreferredEnclaveTypeParameterAsDefault(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().createOrUpdate("sqlcrudtest-2369", "sqlcrudtest-8069",
            "sqlcrudtest-8102", new ElasticPoolInner().withLocation("Japan East")
                .withSku(new Sku().withName("GP_Gen5_4")).withPreferredEnclaveType(AlwaysEncryptedEnclaveType.DEFAULT),
            com.azure.core.util.Context.NONE);
    }
}
