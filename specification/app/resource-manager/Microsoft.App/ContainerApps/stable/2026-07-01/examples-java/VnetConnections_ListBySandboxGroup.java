
/**
 * Samples for VnetConnections ListBySandboxGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/VnetConnections_ListBySandboxGroup.json
     */
    /**
     * Sample code: List VnetConnections by SandboxGroup.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listVnetConnectionsBySandboxGroup(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.vnetConnections().listBySandboxGroup("myRg", "testgroup", com.azure.core.util.Context.NONE);
    }
}
