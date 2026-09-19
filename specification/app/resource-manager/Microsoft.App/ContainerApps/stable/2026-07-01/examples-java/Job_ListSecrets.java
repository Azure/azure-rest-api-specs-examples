
/**
 * Samples for Jobs ListSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Job_ListSecrets.json
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
