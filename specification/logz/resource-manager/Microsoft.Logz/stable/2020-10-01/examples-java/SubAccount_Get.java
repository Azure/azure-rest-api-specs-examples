import com.azure.core.util.Context;

/** Samples for SubAccount Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_Get.json
     */
    /**
     * Sample code: SubAccount_Get.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void subAccountGet(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.subAccounts().getWithResponse("myResourceGroup", "myMonitor", "SubAccount1", Context.NONE);
    }
}
