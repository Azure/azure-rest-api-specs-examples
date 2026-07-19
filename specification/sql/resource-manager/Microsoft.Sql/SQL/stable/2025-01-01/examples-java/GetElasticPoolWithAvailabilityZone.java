
/**
 * Samples for ElasticPools Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetElasticPoolWithAvailabilityZone.json
     */
    /**
     * Sample code: Get an elastic pool with Availability Zone.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAnElasticPoolWithAvailabilityZone(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().getWithResponse("sqlcrudtest-2369", "sqlcrudtest-8069",
            "sqlcrudtest-8102", com.azure.core.util.Context.NONE);
    }
}
