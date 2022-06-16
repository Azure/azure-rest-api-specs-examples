import com.azure.core.util.Context;

/** Samples for ContainerAppsRevisions GetRevision. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetRevision.json
     */
    /**
     * Sample code: Get Container App's revision.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getContainerAppSRevision(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getContainerAppsRevisions()
            .getRevisionWithResponse("rg", "testcontainerApp0", "testcontainerApp0-pjxhsye", Context.NONE);
    }
}
