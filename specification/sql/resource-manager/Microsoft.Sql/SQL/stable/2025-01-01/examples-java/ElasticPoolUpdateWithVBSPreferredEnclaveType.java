
import com.azure.resourcemanager.sql.models.AlwaysEncryptedEnclaveType;
import com.azure.resourcemanager.sql.models.ElasticPoolUpdate;
import com.azure.resourcemanager.sql.models.Sku;

/**
 * Samples for ElasticPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ElasticPoolUpdateWithVBSPreferredEnclaveType.json
     */
    /**
     * Sample code: Update an elastic pool with preferred enclave type parameter as VBS.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAnElasticPoolWithPreferredEnclaveTypeParameterAsVBS(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().update(
            "sqlcrudtest-2369", "sqlcrudtest-8069", "sqlcrudtest-8102", new ElasticPoolUpdate()
                .withSku(new Sku().withName("GP_Gen5_4")).withPreferredEnclaveType(AlwaysEncryptedEnclaveType.VBS),
            com.azure.core.util.Context.NONE);
    }
}
