import com.azure.core.util.Context;

/** Samples for Accounts Head. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/Accounts/Accounts_Head.json
     */
    /**
     * Sample code: Checks whether account exists.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void checksWhetherAccountExists(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.accounts().headWithResponse("test-rg", "contoso", Context.NONE);
    }
}
