/** Samples for ContainerApps List. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/ListContainerAppsBySubscription.json
     */
    /**
     * Sample code: List Container Apps by resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listContainerAppsByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getContainerApps().list(com.azure.core.util.Context.NONE);
    }
}
