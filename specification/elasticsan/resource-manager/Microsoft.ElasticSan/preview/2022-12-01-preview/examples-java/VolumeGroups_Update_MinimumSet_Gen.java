import com.azure.resourcemanager.elasticsan.models.VolumeGroup;

/** Samples for VolumeGroups Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2022-12-01-preview/examples/VolumeGroups_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_Update_MinimumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumeGroupsUpdateMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        VolumeGroup resource =
            manager
                .volumeGroups()
                .getWithResponse(
                    "resourcegroupname", "elasticsanname", "volumegroupname", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
