
/**
 * Samples for Jobs Resume.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Jobs_Resume.json
     */
    /**
     * Sample code: Resume Job.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void resumeJob(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().resume("rg", "testcontainerAppsJob0", com.azure.core.util.Context.NONE);
    }
}
