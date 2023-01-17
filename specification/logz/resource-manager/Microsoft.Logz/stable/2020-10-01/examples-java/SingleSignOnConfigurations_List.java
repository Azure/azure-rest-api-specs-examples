/** Samples for SingleSignOn List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SingleSignOnConfigurations_List.json
     */
    /**
     * Sample code: SingleSignOnConfigurations_List.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void singleSignOnConfigurationsList(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.singleSignOns().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
