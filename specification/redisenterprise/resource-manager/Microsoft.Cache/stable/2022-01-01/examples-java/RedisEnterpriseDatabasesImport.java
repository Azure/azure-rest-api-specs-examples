import com.azure.core.util.Context;
import com.azure.resourcemanager.redisenterprise.models.ImportClusterParameters;
import java.util.Arrays;

/** Samples for Databases ImportMethod. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDatabasesImport.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesImport.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDatabasesImport(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager
            .databases()
            .importMethod(
                "rg1",
                "cache1",
                "default",
                new ImportClusterParameters()
                    .withSasUris(
                        Arrays
                            .asList(
                                "https://contosostorage.blob.core.window.net/urltoBlobFile1?sasKeyParameters",
                                "https://contosostorage.blob.core.window.net/urltoBlobFile2?sasKeyParameters")),
                Context.NONE);
    }
}
