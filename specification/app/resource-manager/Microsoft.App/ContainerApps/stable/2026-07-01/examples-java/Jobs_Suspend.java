
/**
 * Samples for Jobs Suspend.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Jobs_Suspend.json
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
