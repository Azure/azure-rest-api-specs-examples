
/**
 * Samples for ContainerAppsRevisions DeactivateRevision.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Revisions_Deactivate.json
     */
    /**
     * Sample code: Deactivate Container App's revision.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        deactivateContainerAppSRevision(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsRevisions().deactivateRevisionWithResponse("rg", "testcontainerApp0",
            "testcontainerApp0-pjxhsye", com.azure.core.util.Context.NONE);
    }
}
