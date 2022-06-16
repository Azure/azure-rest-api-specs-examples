import com.azure.core.util.Context;

/** Samples for SubAccount Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_Create.json
     */
    /**
     * Sample code: subAccount_Create.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void subAccountCreate(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.subAccounts().create("myResourceGroup", "myMonitor", "SubAccount1", null, Context.NONE);
    }
}
