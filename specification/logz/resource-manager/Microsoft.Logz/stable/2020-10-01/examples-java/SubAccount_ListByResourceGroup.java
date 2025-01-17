
/**
 * Samples for SubAccount List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_ListByResourceGroup.json
     */
    /**
     * Sample code: SubAccount_List.
     * 
     * @param manager Entry point to LogzManager.
     */
    public static void subAccountList(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.subAccounts().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
