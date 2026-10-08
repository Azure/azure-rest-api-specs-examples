
/**
 * Samples for ResourceProvider RestoreVolume.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/RestoreVolume_MinimumSet_Gen.json
     */
    /**
     * Sample code: RestoreVolume_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void restoreVolumeMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.resourceProviders().restoreVolume("resourcegroupname", "elasticsanname", "volumegroupname",
            "volumename-1741526907", com.azure.core.util.Context.NONE);
    }
}
