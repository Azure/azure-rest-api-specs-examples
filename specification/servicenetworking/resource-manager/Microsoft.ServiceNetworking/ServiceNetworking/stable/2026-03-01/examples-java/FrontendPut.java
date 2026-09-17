
import com.azure.resourcemanager.servicenetworking.models.FrontendProperties;
import com.azure.resourcemanager.servicenetworking.models.PublicNetworkAccess;

/**
 * Samples for FrontendsInterface CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/FrontendPut.json
     */
    /**
     * Sample code: Put Frontend.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void putFrontend(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.frontendsInterfaces().define("fe1").withRegion("NorthCentralUS")
            .withExistingTrafficController("rg1", "tc1")
            .withProperties(new FrontendProperties().withPublicNetworkAccess(PublicNetworkAccess.ENABLED)).create();
    }
}
