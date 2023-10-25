import com.azure.resourcemanager.elasticsan.models.XMsDeleteSnapshots;
import com.azure.resourcemanager.elasticsan.models.XMsForceDelete;

/** Samples for Volumes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Volumes_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumesDeleteMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager
            .volumes()
            .delete(
                "resourcegroupname",
                "elasticsanname",
                "volumegroupname",
                "volumename",
                XMsDeleteSnapshots.TRUE,
                XMsForceDelete.TRUE,
                com.azure.core.util.Context.NONE);
    }
}
