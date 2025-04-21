
/**
 * Samples for VolumeGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/
     * VolumeGroups_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumeGroupsGetMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeGroups().getWithResponse("resourcegroupname", "elasticsanname", "volumegroupname",
            com.azure.core.util.Context.NONE);
    }
}
