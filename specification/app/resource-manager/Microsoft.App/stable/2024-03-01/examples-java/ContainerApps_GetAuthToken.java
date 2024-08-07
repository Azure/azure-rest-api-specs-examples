
/**
 * Samples for ContainerApps GetAuthToken.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ContainerApps_GetAuthToken.json
     */
    /**
     * Sample code: Get Container App Auth Token.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getContainerAppAuthToken(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().getAuthTokenWithResponse("rg", "testcontainerapp0", com.azure.core.util.Context.NONE);
    }
}
