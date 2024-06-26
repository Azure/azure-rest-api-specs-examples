
/**
 * Samples for ContainerApps ListSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ListContainerAppSecrets.json
     */
    /**
     * Sample code: List Container Apps Secrets.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listContainerAppsSecrets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getContainerApps().listSecretsWithResponse("testcontainerApp0",
            com.azure.core.util.Context.NONE);
    }
}
