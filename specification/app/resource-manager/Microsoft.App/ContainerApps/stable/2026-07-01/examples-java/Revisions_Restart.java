
/**
 * Samples for ContainerAppsRevisions RestartRevision.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Revisions_Restart.json
     */
    /**
     * Sample code: Restart Container App's revision.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        restartContainerAppSRevision(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsRevisions().restartRevisionWithResponse("rg", "testStaticSite0",
            "testcontainerApp0-pjxhsye", com.azure.core.util.Context.NONE);
    }
}
