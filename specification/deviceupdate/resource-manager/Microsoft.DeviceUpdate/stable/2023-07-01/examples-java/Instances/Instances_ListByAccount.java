
/**
 * Samples for Instances ListByAccount.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/Instances/
     * Instances_ListByAccount.json
     */
    /**
     * Sample code: Gets list of Instances by Account.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void
        getsListOfInstancesByAccount(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.instances().listByAccount("test-rg", "contoso", com.azure.core.util.Context.NONE);
    }
}
