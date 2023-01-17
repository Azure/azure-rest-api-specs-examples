/** Samples for SubAccount Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_Delete.json
     */
    /**
     * Sample code: SubAccount_Delete.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void subAccountDelete(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.subAccounts().delete("myResourceGroup", "myMonitor", "someName", com.azure.core.util.Context.NONE);
    }
}
