
/**
 * Samples for ConnectedEnvironmentsDaprComponents ListSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ConnectedEnvironmentsDaprComponents_ListSecrets.json
     */
    /**
     * Sample code: List Container Apps Secrets.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listContainerAppsSecrets(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironmentsDaprComponents().listSecretsWithResponse("examplerg", "myenvironment", "reddog",
            com.azure.core.util.Context.NONE);
    }
}
