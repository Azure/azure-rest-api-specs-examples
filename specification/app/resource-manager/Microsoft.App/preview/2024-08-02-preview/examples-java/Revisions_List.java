
/**
 * Samples for ContainerAppsDiagnostics ListRevisions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/Revisions_List.json
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
