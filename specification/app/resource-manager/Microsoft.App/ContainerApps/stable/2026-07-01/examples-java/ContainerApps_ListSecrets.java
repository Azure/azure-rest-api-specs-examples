
/**
 * Samples for ContainerApps ListSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerApps_ListSecrets.json
     */
    /**
     * Sample code: List Container Apps Secrets.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listContainerAppsSecrets(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().listSecretsWithResponse("rg", "testcontainerApp0", com.azure.core.util.Context.NONE);
    }
}
