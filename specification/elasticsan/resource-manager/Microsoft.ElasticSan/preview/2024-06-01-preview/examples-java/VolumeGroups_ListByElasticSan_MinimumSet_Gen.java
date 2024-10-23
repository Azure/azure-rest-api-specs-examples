
/**
 * Samples for VolumeGroups ListByElasticSan.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/
     * VolumeGroups_ListByElasticSan_MinimumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_ListByElasticSan_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        volumeGroupsListByElasticSanMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeGroups().listByElasticSan("resourcegroupname", "elasticsanname",
            com.azure.core.util.Context.NONE);
    }
}
