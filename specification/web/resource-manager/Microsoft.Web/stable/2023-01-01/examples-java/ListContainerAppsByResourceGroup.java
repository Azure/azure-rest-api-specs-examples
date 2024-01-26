
/**
 * Samples for ContainerApps ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/ListContainerAppsByResourceGroup.json
     */
    /**
     * Sample code: List Container Apps by resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listContainerAppsByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getContainerApps().listByResourceGroup("rg",
            com.azure.core.util.Context.NONE);
    }
}
