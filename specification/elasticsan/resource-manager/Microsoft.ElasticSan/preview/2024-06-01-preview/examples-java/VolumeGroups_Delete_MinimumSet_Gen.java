
/**
 * Samples for VolumeGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/
     * VolumeGroups_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumeGroupsDeleteMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeGroups().delete("resourcegroupname", "elasticsanname", "volumegroupname",
            com.azure.core.util.Context.NONE);
    }
}
