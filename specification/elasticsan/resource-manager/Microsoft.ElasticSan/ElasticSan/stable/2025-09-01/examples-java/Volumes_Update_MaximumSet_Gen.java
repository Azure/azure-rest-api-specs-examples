
import com.azure.resourcemanager.elasticsan.models.ManagedByInfo;
import com.azure.resourcemanager.elasticsan.models.Volume;

/**
 * Samples for Volumes Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Volumes_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Volumes_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumesUpdateMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        Volume resource = manager.volumes().getWithResponse("resourcegroupname", "elasticsanname", "volumegroupname",
            "volumename", com.azure.core.util.Context.NONE).getValue();
        resource.update().withSizeGiB(7L)
            .withManagedBy(new ManagedByInfo().withResourceId("pclpkrpkpmvcsegcubrakcoodrubo")).apply();
    }
}
