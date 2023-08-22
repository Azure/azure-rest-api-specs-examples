/** Samples for ContainerAppsRevisions ListRevisions. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Revisions_List.json
     */
    /**
     * Sample code: List Container App's revisions.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listContainerAppSRevisions(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsRevisions()
            .listRevisions("rg", "testcontainerApp0", null, com.azure.core.util.Context.NONE);
    }
}
