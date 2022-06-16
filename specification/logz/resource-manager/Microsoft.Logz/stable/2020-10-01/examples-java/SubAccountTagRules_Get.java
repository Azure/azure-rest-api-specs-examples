import com.azure.core.util.Context;

/** Samples for SubAccountTagRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccountTagRules_Get.json
     */
    /**
     * Sample code: SubAccountTagRules_Get.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void subAccountTagRulesGet(com.azure.resourcemanager.logz.LogzManager manager) {
        manager
            .subAccountTagRules()
            .getWithResponse("myResourceGroup", "myMonitor", "SubAccount1", "default", Context.NONE);
    }
}
