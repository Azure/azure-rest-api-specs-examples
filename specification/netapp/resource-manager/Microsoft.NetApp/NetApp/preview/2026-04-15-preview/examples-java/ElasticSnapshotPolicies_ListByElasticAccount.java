
/**
 * Samples for ElasticSnapshotPolicies ListByElasticAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-15-preview/ElasticSnapshotPolicies_ListByElasticAccount.json
     */
    /**
     * Sample code: ElasticSnapshotPolicies_ListByElasticAccount.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        elasticSnapshotPoliciesListByElasticAccount(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticSnapshotPolicies().listByElasticAccount("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
