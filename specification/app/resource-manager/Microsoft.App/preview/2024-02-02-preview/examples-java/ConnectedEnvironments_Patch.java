
/**
 * Samples for ConnectedEnvironments Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/ConnectedEnvironments_Patch.
     * json
     */
    /**
     * Sample code: Patch Managed Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        patchManagedEnvironment(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironments().updateWithResponse("examplerg", "testenv", com.azure.core.util.Context.NONE);
    }
}
