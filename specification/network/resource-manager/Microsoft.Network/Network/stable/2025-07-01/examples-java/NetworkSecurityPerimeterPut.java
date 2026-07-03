
import com.azure.resourcemanager.network.fluent.models.NetworkSecurityPerimeterInner;

/**
 * Samples for NetworkSecurityPerimeters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkSecurityPerimeterPut.json
     */
    /**
     * Sample code: Put Network Security Perimeter.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void putNetworkSecurityPerimeter(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeters().createOrUpdateWithResponse("rg1", "nsp1",
            new NetworkSecurityPerimeterInner().withLocation("location1"), com.azure.core.util.Context.NONE);
    }
}
