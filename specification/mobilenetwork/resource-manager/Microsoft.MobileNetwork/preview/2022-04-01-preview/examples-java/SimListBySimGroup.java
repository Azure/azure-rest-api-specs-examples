import com.azure.core.util.Context;

/** Samples for Sims ListBySimGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimListBySimGroup.json
     */
    /**
     * Sample code: List SIMs in a SIM group.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listSIMsInASIMGroup(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.sims().listBySimGroup("rg1", "testSimGroup", Context.NONE);
    }
}
