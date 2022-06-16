import com.azure.core.util.Context;

/** Samples for Sims GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SimGet.json
     */
    /**
     * Sample code: Get sim.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getSim(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.sims().getByResourceGroupWithResponse("testResourceGroupName", "testSimName", Context.NONE);
    }
}
