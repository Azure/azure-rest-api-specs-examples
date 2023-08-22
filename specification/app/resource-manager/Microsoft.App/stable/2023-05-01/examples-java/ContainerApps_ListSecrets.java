/** Samples for ContainerApps ListSecrets. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ContainerApps_ListSecrets.json
     */
    /**
     * Sample code: List Container Apps Secrets.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listContainerAppsSecrets(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().listSecretsWithResponse("rg", "testcontainerApp0", com.azure.core.util.Context.NONE);
    }
}
