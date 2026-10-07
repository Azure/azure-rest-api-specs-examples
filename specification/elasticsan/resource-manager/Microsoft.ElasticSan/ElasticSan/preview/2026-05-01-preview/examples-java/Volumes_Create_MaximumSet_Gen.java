
import com.azure.resourcemanager.elasticsan.models.ManagedByResources;
import com.azure.resourcemanager.elasticsan.models.SourceCreationData;
import com.azure.resourcemanager.elasticsan.models.VolumeCreateOption;
import java.util.Arrays;

/**
 * Samples for Volumes Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Volumes_Create_MaximumSet_Gen.json
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
            .withManagedBy(Arrays.asList(new ManagedByResources().withClientId("pclpkrpkpmvcsegcubrakcoodrubo")
                .withVersion(1)
                .withResourceIds(Arrays.asList(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.SomeProvider/someResource/myResource"))))
            .create();
    }
}
