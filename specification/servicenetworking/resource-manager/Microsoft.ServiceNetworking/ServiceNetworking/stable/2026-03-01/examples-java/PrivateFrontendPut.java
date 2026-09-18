
import com.azure.resourcemanager.servicenetworking.models.FrontendAssociation;
import com.azure.resourcemanager.servicenetworking.models.FrontendProperties;
import com.azure.resourcemanager.servicenetworking.models.PublicNetworkAccess;

/**
 * Samples for FrontendsInterface CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PrivateFrontendPut.json
     */
    /**
     * Sample code: Put Private Frontend.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void
        putPrivateFrontend(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.frontendsInterfaces().define("pfe1").withRegion("NorthCentralUS")
            .withExistingTrafficController("rg1", "tc1")
            .withProperties(new FrontendProperties().withPublicNetworkAccess(PublicNetworkAccess.DISABLED)
                .withAssociation(new FrontendAssociation().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/tc1/associations/as1")))
            .create();
    }
}
