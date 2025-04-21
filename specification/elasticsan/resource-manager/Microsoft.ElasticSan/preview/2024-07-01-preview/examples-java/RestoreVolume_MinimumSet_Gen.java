
/**
 * Samples for ResourceProvider RestoreVolume.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/
     * RestoreVolume_MinimumSet_Gen.json
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
