
/**
 * Samples for BuildAuthToken List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/Builds_ListAuthToken.json
     */
    /**
     * Sample code: Get Build Auth Token.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getBuildAuthToken(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.buildAuthTokens().listWithResponse("rg", "testBuilder", "testBuild", com.azure.core.util.Context.NONE);
    }
}
