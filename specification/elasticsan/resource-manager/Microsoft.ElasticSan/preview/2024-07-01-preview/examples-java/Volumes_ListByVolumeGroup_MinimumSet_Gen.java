
import com.azure.resourcemanager.elasticsan.models.XMsAccessSoftDeletedResources;

/**
 * Samples for Volumes ListByVolumeGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/
     * Volumes_ListByVolumeGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: Volumes_ListByVolumeGroup_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        volumesListByVolumeGroupMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumes().listByVolumeGroup("resourcegroupname", "elasticsanname", "volumegroupname",
            XMsAccessSoftDeletedResources.TRUE, com.azure.core.util.Context.NONE);
    }
}
