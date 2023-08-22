/** Samples for ConnectedEnvironmentsDaprComponents ListSecrets. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ConnectedEnvironmentsDaprComponents_ListSecrets.json
     */
    /**
     * Sample code: List Container Apps Secrets.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listContainerAppsSecrets(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .connectedEnvironmentsDaprComponents()
            .listSecretsWithResponse("examplerg", "myenvironment", "reddog", com.azure.core.util.Context.NONE);
    }
}
