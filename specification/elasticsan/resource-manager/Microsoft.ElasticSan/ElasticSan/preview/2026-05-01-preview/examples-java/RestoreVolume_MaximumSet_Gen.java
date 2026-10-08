
/**
 * Samples for ResourceProvider RestoreVolume.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/RestoreVolume_MaximumSet_Gen.json
     */
    /**
     * Sample code: RestoreVolume_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void restoreVolumeMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.resourceProviders().restoreVolume("resourcegroupname", "elasticsanname", "volumegroupname",
            "volumename-1741526907", com.azure.core.util.Context.NONE);
    }
}
