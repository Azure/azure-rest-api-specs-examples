
import com.azure.resourcemanager.redisenterprise.models.ExportClusterParameters;

/**
 * Samples for Databases Export.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/
     * RedisEnterpriseDatabasesExport.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesExport.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void
        redisEnterpriseDatabasesExport(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.databases().export("rg1", "cache1", "default",
            new ExportClusterParameters()
                .withSasUri("https://contosostorage.blob.core.window.net/urlToBlobContainer?sasKeyParameters"),
            com.azure.core.util.Context.NONE);
    }
}
