
/**
 * Samples for Jobs Suspend.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/Jobs_Suspend.
     * json
     */
    /**
     * Sample code: Suspend Job.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void suspendJob(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().suspend("rg", "testcontainerAppsJob0", com.azure.core.util.Context.NONE);
    }
}
