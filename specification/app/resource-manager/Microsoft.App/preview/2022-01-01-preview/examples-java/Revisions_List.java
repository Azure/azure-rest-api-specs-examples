import com.azure.core.util.Context;

/** Samples for ContainerAppsRevisions ListRevisions. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/Revisions_List.json
     */
    /**
     * Sample code: List Container App's revisions.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listContainerAppSRevisions(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsRevisions().listRevisions("rg", "testcontainerApp0", Context.NONE);
    }
}
