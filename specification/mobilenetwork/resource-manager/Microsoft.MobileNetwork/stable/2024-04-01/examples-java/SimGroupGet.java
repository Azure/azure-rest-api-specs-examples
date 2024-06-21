
/**
 * Samples for SimGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/SimGroupGet.json
     */
    /**
     * Sample code: Get SIM group.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getSIMGroup(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simGroups().getByResourceGroupWithResponse("testResourceGroupName", "testSimGroupName",
            com.azure.core.util.Context.NONE);
    }
}
