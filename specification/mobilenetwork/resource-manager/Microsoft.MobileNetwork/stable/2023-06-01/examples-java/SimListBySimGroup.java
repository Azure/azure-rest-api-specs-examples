/** Samples for Sims ListByGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SimListBySimGroup.json
     */
    /**
     * Sample code: List SIMs in a SIM group.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listSIMsInASIMGroup(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.sims().listByGroup("rg1", "testSimGroup", com.azure.core.util.Context.NONE);
    }
}
