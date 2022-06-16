import com.azure.core.util.Context;

/** Samples for ContainerApps ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListContainerAppsByResourceGroup.json
     */
    /**
     * Sample code: List Container Apps by resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listContainerAppsByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getContainerApps().listByResourceGroup("rg", Context.NONE);
    }
}
