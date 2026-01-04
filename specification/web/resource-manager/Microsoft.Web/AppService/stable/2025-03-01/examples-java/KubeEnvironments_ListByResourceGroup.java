
/**
 * Samples for KubeEnvironments ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * KubeEnvironments_ListByResourceGroup.json
     */
    /**
     * Sample code: List kube environments by resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listKubeEnvironmentsByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getKubeEnvironments().listByResourceGroup("examplerg",
            com.azure.core.util.Context.NONE);
    }
}
