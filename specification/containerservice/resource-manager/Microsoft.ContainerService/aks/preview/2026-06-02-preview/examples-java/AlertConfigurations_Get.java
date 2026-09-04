
/**
 * Samples for AlertConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/AlertConfigurations_Get.json
     */
    /**
     * Sample code: Get Alert Configuration.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getAlertConfiguration(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAlertConfigurations().getWithResponse("rg1", "clustername1", "alertconfig1",
            com.azure.core.util.Context.NONE);
    }
}
