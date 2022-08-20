import com.azure.core.util.Context;

/** Samples for Accounts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/Accounts/Accounts_Get.json
     */
    /**
     * Sample code: Gets Account details.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void getsAccountDetails(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.accounts().getByResourceGroupWithResponse("test-rg", "contoso", Context.NONE);
    }
}
