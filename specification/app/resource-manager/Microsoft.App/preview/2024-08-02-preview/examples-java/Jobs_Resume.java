
/**
 * Samples for Jobs Resume.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/Jobs_Resume.json
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
