import com.azure.core.util.Context;

/** Samples for Instances ListByAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/Instances/Instances_ListByAccount.json
     */
    /**
     * Sample code: Gets list of Instances by Account.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void getsListOfInstancesByAccount(
        com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.instances().listByAccount("test-rg", "contoso", Context.NONE);
    }
}
