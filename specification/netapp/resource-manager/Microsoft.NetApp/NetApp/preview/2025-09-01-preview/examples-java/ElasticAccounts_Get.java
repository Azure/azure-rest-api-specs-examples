
/**
 * Samples for ElasticAccounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticAccounts_Get.json
     */
    /**
     * Sample code: ElasticAccounts_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticAccountsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticAccounts().getByResourceGroupWithResponse("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
