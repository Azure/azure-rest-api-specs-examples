
/**
 * Samples for KubeEnvironments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/KubeEnvironments_Delete.
     * json
     */
    /**
     * Sample code: Delete kube environment by name.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteKubeEnvironmentByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getKubeEnvironments().delete("examplerg", "examplekenv",
            com.azure.core.util.Context.NONE);
    }
}
