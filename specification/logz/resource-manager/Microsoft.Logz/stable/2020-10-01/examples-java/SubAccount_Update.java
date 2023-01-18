/** Samples for SubAccount Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_Update.json
     */
    /**
     * Sample code: SubAccount_Update.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void subAccountUpdate(com.azure.resourcemanager.logz.LogzManager manager) {
        manager
            .subAccounts()
            .updateWithResponse("myResourceGroup", "myMonitor", "SubAccount1", null, com.azure.core.util.Context.NONE);
    }
}
