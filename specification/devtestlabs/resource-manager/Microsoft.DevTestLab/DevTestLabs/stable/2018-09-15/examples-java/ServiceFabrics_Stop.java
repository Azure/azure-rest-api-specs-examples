
/**
 * Samples for ServiceFabrics Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/ServiceFabrics_Stop.json
     */
    /**
     * Sample code: ServiceFabrics_Stop.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricsStop(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.serviceFabrics().stop("resourceGroupName", "{labName}", "{userName}", "{serviceFabricName}",
            com.azure.core.util.Context.NONE);
    }
}
