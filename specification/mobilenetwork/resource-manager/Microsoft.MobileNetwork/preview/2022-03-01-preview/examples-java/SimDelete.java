import com.azure.core.util.Context;

/** Samples for Sims Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimDelete.json
     */
    /**
     * Sample code: Delete sim.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteSim(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.sims().delete("testResourceGroupName", "testSim", Context.NONE);
    }
}
