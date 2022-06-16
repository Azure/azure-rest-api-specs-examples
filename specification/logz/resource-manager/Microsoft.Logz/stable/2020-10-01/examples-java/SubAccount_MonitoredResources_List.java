import com.azure.core.util.Context;

/** Samples for SubAccount ListMonitoredResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_MonitoredResources_List.json
     */
    /**
     * Sample code: SubAccount_MonitoredResources_List.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void subAccountMonitoredResourcesList(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.subAccounts().listMonitoredResources("myResourceGroup", "myMonitor", "SubAccount1", Context.NONE);
    }
}
