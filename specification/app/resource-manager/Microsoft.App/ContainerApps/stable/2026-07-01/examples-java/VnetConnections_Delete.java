
/**
 * Samples for VnetConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/VnetConnections_Delete.json
     */
    /**
     * Sample code: Delete a VnetConnection.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteAVnetConnection(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.vnetConnections().delete("myRg", "testgroup", "myVnetConnection", com.azure.core.util.Context.NONE);
    }
}
