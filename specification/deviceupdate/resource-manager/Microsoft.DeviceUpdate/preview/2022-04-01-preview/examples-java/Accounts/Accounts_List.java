import com.azure.core.util.Context;

/** Samples for Accounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/Accounts/Accounts_List.json
     */
    /**
     * Sample code: Get list of Accounts.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void getListOfAccounts(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.accounts().list(Context.NONE);
    }
}
