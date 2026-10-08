
import com.azure.resourcemanager.elasticsan.models.Volume;

/**
 * Samples for Volumes Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Volumes_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: Volumes_Update_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumesUpdateMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        Volume resource = manager.volumes().getWithResponse("resourcegroupname", "elasticsanname", "volumegroupname",
            "volumename", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
