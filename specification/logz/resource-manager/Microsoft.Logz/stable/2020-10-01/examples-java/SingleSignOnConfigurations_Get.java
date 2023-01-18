/** Samples for SingleSignOn Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SingleSignOnConfigurations_Get.json
     */
    /**
     * Sample code: SingleSignOnConfigurations_Get.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void singleSignOnConfigurationsGet(com.azure.resourcemanager.logz.LogzManager manager) {
        manager
            .singleSignOns()
            .getWithResponse("myResourceGroup", "myMonitor", "default", com.azure.core.util.Context.NONE);
    }
}
