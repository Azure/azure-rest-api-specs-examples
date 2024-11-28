
/**
 * Samples for ContainerAppsRevisions DeactivateRevision.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/DeactivateRevision.json
     */
    /**
     * Sample code: Deactivate Container App's revision.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deactivateContainerAppSRevision(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getContainerAppsRevisions().deactivateRevisionWithResponse("rg",
            "testcontainerApp0", "testcontainerApp0-pjxhsye", com.azure.core.util.Context.NONE);
    }
}
