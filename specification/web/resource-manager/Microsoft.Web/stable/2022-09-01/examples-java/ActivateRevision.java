/** Samples for ContainerAppsRevisions ActivateRevision. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/ActivateRevision.json
     */
    /**
     * Sample code: Activate Container App's revision.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void activateContainerAppSRevision(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getContainerAppsRevisions()
            .activateRevisionWithResponse(
                "rg", "testcontainerApp0", "testcontainerApp0-pjxhsye", com.azure.core.util.Context.NONE);
    }
}
