import com.azure.core.util.Context;

/** Samples for Accounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/Accounts/Accounts_Delete.json
     */
    /**
     * Sample code: Deletes an account.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void deletesAnAccount(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.accounts().delete("test-rg", "contoso", Context.NONE);
    }
}
