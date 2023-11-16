/** Samples for Sims Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/SimGet.json
     */
    /**
     * Sample code: Get SIM.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getSIM(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .sims()
            .getWithResponse("testResourceGroupName", "testSimGroup", "testSimName", com.azure.core.util.Context.NONE);
    }
}
