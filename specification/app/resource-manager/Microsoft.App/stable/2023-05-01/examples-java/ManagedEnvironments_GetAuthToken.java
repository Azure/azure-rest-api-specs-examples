/** Samples for ManagedEnvironments GetAuthToken. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ManagedEnvironments_GetAuthToken.json
     */
    /**
     * Sample code: Get Managed Environment Auth Token.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getManagedEnvironmentAuthToken(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironments().getAuthTokenWithResponse("rg", "testenv", com.azure.core.util.Context.NONE);
    }
}
