
/**
 * Samples for ElasticSnapshotPolicies ListElasticVolumes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-15-preview/ElasticSnapshotPolicies_ListElasticVolumes.json
     */
    /**
     * Sample code: ElasticSnapshotPolicies_ListElasticVolumes.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        elasticSnapshotPoliciesListElasticVolumes(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticSnapshotPolicies().listElasticVolumes("myRG", "account1", "snapshotPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
