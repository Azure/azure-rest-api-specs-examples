
/**
 * Samples for ContainerAppsRevisions RestartRevision.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/RestartRevision.json
     */
    /**
     * Sample code: Restart Container App's revision.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void restartContainerAppSRevision(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getContainerAppsRevisions().restartRevisionWithResponse("rg",
            "testStaticSite0", "testcontainerApp0-pjxhsye", com.azure.core.util.Context.NONE);
    }
}
