
/**
 * Samples for ContainerApps GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/GetContainerApp.json
     */
    /**
     * Sample code: Get Container App.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getContainerApp(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getContainerApps().getByResourceGroupWithResponse("rg",
            "testcontainerApp0", com.azure.core.util.Context.NONE);
    }
}
