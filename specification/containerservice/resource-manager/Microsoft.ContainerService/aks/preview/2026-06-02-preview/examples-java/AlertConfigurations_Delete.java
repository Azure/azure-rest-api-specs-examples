
/**
 * Samples for AlertConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/AlertConfigurations_Delete.json
     */
    /**
     * Sample code: Delete Alert Configuration.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        deleteAlertConfiguration(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAlertConfigurations().delete("rg1", "clustername1", "alertconfig1",
            com.azure.core.util.Context.NONE);
    }
}
