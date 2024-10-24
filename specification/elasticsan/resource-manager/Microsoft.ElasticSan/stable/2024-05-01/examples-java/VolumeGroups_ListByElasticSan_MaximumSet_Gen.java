
/**
 * Samples for VolumeGroups ListByElasticSan.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2024-05-01/examples/
     * VolumeGroups_ListByElasticSan_MaximumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_ListByElasticSan_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        volumeGroupsListByElasticSanMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeGroups().listByElasticSan("resourcegroupname", "elasticsanname",
            com.azure.core.util.Context.NONE);
    }
}
