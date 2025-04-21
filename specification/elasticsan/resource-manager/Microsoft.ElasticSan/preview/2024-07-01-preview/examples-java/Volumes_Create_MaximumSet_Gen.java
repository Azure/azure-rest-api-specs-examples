
import com.azure.resourcemanager.elasticsan.models.ManagedByInfo;
import com.azure.resourcemanager.elasticsan.models.SourceCreationData;
import com.azure.resourcemanager.elasticsan.models.VolumeCreateOption;

/**
 * Samples for Volumes Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/
     * Volumes_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: Volumes_Create_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumesCreateMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumes().define("volumename")
            .withExistingVolumegroup("resourcegroupname", "elasticsanname", "volumegroupname").withSizeGiB(23L)
            .withCreationData(
                new SourceCreationData().withCreateSource(VolumeCreateOption.NONE).withSourceId("mdonegivjquite"))
            .withManagedBy(new ManagedByInfo().withResourceId("pclpkrpkpmvcsegcubrakcoodrubo")).create();
    }
}
