
/**
 * Samples for ElasticAccounts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-15-preview/ElasticAccounts_ListBySubscription.json
     */
    /**
     * Sample code: ElasticAccounts_ListBySubscription.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticAccountsListBySubscription(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticAccounts().list(com.azure.core.util.Context.NONE);
    }
}
