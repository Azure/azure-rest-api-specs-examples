
import com.azure.resourcemanager.netapp.models.Exclude;
import com.azure.resourcemanager.netapp.models.ListReplicationsRequest;

/**
 * Samples for Volumes ListReplications.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Volumes_ListReplications.json
     */
    /**
     * Sample code: Volumes_ListReplications.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesListReplications(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().listReplications("myRG", "account1", "pool1", "volume1",
            new ListReplicationsRequest().withExclude(Exclude.NONE), com.azure.core.util.Context.NONE);
    }
}
