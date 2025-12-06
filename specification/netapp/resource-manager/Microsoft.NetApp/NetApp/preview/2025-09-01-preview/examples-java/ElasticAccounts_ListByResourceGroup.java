
/**
 * Samples for ElasticAccounts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticAccounts_ListByResourceGroup.json
     */
    /**
     * Sample code: ElasticAccounts_ListByResourceGroup.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticAccountsListByResourceGroup(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticAccounts().listByResourceGroup("myRG", com.azure.core.util.Context.NONE);
    }
}
