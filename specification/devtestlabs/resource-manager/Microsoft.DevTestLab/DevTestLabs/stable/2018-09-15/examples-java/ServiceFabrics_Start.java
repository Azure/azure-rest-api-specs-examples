
/**
 * Samples for ServiceFabrics Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/ServiceFabrics_Start.json
     */
    /**
     * Sample code: ServiceFabrics_Start.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricsStart(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.serviceFabrics().start("resourceGroupName", "{labName}", "{userName}", "{serviceFabricName}",
            com.azure.core.util.Context.NONE);
    }
}
