
/**
 * Samples for Jobs ListSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/Job_ListSecrets.json
     */
    /**
     * Sample code: List Container Apps Job Secrets.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listContainerAppsJobSecrets(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().listSecretsWithResponse("rg", "testcontainerAppsJob0", com.azure.core.util.Context.NONE);
    }
}
