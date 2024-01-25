
import com.azure.resourcemanager.elasticsan.models.Volume;

/**
 * Samples for Volumes Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/
     * Volumes_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Volumes_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumesUpdateMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        Volume resource = manager.volumes().getWithResponse("resourcegroupname", "elasticsanname", "volumegroupname",
            "volumename", com.azure.core.util.Context.NONE).getValue();
        resource.update().withSizeGiB(11L).apply();
    }
}
