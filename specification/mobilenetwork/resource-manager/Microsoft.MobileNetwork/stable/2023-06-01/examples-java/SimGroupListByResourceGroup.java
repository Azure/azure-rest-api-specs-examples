/** Samples for SimGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SimGroupListByResourceGroup.json
     */
    /**
     * Sample code: List SIM groups in a resource group.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listSIMGroupsInAResourceGroup(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simGroups().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
