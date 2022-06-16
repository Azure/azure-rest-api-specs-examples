import com.azure.core.util.Context;

/** Samples for SubAccountTagRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccountTagRules_CreateOrUpdate.json
     */
    /**
     * Sample code: SubAccountTagRules_CreateOrUpdate.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void subAccountTagRulesCreateOrUpdate(com.azure.resourcemanager.logz.LogzManager manager) {
        manager
            .subAccountTagRules()
            .createOrUpdateWithResponse("myResourceGroup", "myMonitor", "SubAccount1", "default", null, Context.NONE);
    }
}
