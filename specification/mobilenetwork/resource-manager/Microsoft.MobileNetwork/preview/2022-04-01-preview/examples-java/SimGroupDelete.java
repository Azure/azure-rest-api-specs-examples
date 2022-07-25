import com.azure.core.util.Context;

/** Samples for SimGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimGroupDelete.json
     */
    /**
     * Sample code: Delete SIM group.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteSIMGroup(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simGroups().delete("testResourceGroupName", "testSimGroup", Context.NONE);
    }
}
