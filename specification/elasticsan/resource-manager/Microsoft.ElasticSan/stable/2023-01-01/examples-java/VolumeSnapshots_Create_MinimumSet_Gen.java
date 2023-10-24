import com.azure.resourcemanager.elasticsan.models.SnapshotCreationData;

/** Samples for VolumeSnapshots Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_Create_MinimumSet_Gen.json
     */
    /**
     * Sample code: VolumeSnapshots_Create_MinimumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumeSnapshotsCreateMinimumSetGen(
        com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager
            .volumeSnapshots()
            .define("snapshotname")
            .withExistingVolumegroup("resourcegroupname", "elasticsanname", "volumegroupname")
            .withCreationData(
                new SnapshotCreationData()
                    .withSourceId(
                        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"))
            .create();
    }
}
