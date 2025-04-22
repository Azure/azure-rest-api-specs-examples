
import com.azure.resourcemanager.elasticsan.models.DeleteType;

/**
 * Samples for Volumes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/
     * Volumes_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Volumes_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumesDeleteMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumes().delete("resourcegroupname", "elasticsanname", "volumegroupname", "volumename", null, null,
            DeleteType.PERMANENT, com.azure.core.util.Context.NONE);
    }
}
