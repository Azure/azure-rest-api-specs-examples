/** Samples for SimGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SimGroupDelete.json
     */
    /**
     * Sample code: Delete SIM group.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteSIMGroup(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simGroups().delete("testResourceGroupName", "testSimGroup", com.azure.core.util.Context.NONE);
    }
}
