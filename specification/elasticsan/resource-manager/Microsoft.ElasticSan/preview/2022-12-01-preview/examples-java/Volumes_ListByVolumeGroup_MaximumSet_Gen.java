/** Samples for Volumes ListByVolumeGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2022-12-01-preview/examples/Volumes_ListByVolumeGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Volumes_ListByVolumeGroup_MaximumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumesListByVolumeGroupMaximumSetGen(
        com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager
            .volumes()
            .listByVolumeGroup(
                "resourcegroupname", "elasticsanname", "volumegroupname", com.azure.core.util.Context.NONE);
    }
}
