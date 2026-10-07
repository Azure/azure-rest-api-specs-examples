
import com.azure.resourcemanager.elasticsan.models.XMsAccessSoftDeletedResources;

/**
 * Samples for VolumeGroups ListByElasticSan.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/VolumeGroups_ListByElasticSan_MinimumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_ListByElasticSan_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        volumeGroupsListByElasticSanMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeGroups().listByElasticSan("resourcegroupname", "elasticsanname",
            XMsAccessSoftDeletedResources.TRUE, com.azure.core.util.Context.NONE);
    }
}
