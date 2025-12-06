
/**
 * Samples for ElasticAccounts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticAccounts_Delete.json
     */
    /**
     * Sample code: ElasticAccounts_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticAccountsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticAccounts().delete("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
