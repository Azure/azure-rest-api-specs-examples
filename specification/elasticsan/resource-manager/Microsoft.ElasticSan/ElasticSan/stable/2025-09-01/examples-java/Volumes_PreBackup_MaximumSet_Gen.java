
import com.azure.resourcemanager.elasticsan.models.VolumeNameList;
import java.util.Arrays;

/**
 * Samples for Volumes PreBackup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Volumes_PreBackup_MaximumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_PreBackup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        volumeGroupsPreBackupMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumes().preBackup("resourcegroupname", "elasticsanname", "volumegroupname",
            new VolumeNameList().withVolumeNames(Arrays.asList("volumename")), com.azure.core.util.Context.NONE);
    }
}
