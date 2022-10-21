import com.azure.core.util.Context;

/** Samples for Volumes ListByVolumeGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/Volumes_ListByVolumeGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Volumes_ListByVolumeGroup_MaximumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumesListByVolumeGroupMaximumSetGen(
        com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumes().listByVolumeGroup("rgelasticsan", "ti7q-k952-1qB3J_5", "u_5I_1j4t3", Context.NONE);
    }
}
