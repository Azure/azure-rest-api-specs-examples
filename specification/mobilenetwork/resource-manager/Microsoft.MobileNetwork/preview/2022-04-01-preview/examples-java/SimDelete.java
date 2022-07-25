import com.azure.core.util.Context;

/** Samples for Sims Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimDelete.json
     */
    /**
     * Sample code: Delete SIM.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteSIM(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.sims().delete("testResourceGroupName", "testSimGroup", "testSim", Context.NONE);
    }
}
