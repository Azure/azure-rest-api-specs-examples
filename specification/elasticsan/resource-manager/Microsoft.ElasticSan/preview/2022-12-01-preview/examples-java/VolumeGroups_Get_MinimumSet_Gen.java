/** Samples for VolumeGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2022-12-01-preview/examples/VolumeGroups_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_Get_MinimumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumeGroupsGetMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager
            .volumeGroups()
            .getWithResponse(
                "resourcegroupname", "elasticsanname", "volumegroupname", com.azure.core.util.Context.NONE);
    }
}
