
/**
 * Samples for ContainerAppsDiagnostics ListRevisions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Revisions_List1.json
     */
    /**
     * Sample code: List Container App's revisions.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listContainerAppSRevisions(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsDiagnostics().listRevisions("rg", "testcontainerApp0", null,
            com.azure.core.util.Context.NONE);
    }
}
