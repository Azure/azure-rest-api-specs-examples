
/**
 * Samples for ContainerAppsRevisions ListRevisions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/ListRevisions.json
     */
    /**
     * Sample code: List Container App's revisions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listContainerAppSRevisions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getContainerAppsRevisions().listRevisions("rg", "testcontainerApp0",
            com.azure.core.util.Context.NONE);
    }
}
