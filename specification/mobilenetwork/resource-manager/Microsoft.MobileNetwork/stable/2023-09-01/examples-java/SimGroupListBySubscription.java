/** Samples for SimGroups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/SimGroupListBySubscription.json
     */
    /**
     * Sample code: List SIM groups in a subscription.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listSIMGroupsInASubscription(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simGroups().list(com.azure.core.util.Context.NONE);
    }
}
