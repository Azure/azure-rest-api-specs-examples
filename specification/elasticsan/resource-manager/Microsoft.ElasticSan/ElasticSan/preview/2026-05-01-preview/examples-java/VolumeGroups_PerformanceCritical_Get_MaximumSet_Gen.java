
/**
 * Samples for VolumeGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/VolumeGroups_PerformanceCritical_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_PerformanceCritical_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumeGroupsPerformanceCriticalGetMaximumSetGen(
        com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeGroups().getWithResponse("resourcegroupname", "elasticsanname", "volumegroupname",
            com.azure.core.util.Context.NONE);
    }
}
