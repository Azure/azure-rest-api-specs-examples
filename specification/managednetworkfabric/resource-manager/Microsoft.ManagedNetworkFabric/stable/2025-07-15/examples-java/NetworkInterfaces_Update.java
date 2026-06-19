
import com.azure.resourcemanager.managednetworkfabric.models.NetworkInterface;

/**
 * Samples for NetworkInterfaces Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkInterfaces_Update.json
     */
    /**
     * Sample code: NetworkInterfaces_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkInterfacesUpdateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        NetworkInterface resource = manager.networkInterfaces()
            .getWithResponse("example-rg", "example-device", "example-interface", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withAdditionalDescription("device 1").withAnnotation("annotation").apply();
    }
}
