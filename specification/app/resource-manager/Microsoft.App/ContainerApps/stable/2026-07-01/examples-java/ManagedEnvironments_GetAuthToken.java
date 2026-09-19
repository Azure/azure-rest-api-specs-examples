
/**
 * Samples for ManagedEnvironments GetAuthToken.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironments_GetAuthToken.json
     */
    /**
     * Sample code: Get Managed Environment Auth Token.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getManagedEnvironmentAuthToken(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironments().getAuthTokenWithResponse("rg", "testenv", com.azure.core.util.Context.NONE);
    }
}
