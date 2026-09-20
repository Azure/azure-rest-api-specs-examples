
/**
 * Samples for ContainerApps GetAuthToken.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerApps_GetAuthToken.json
     */
    /**
     * Sample code: Get Container App Auth Token.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getContainerAppAuthToken(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().getAuthTokenWithResponse("rg", "testcontainerApp0", com.azure.core.util.Context.NONE);
    }
}
