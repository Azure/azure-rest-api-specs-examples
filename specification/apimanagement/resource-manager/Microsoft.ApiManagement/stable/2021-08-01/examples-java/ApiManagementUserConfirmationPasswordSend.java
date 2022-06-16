import com.azure.core.util.Context;

/** Samples for UserConfirmationPassword Send. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUserConfirmationPasswordSend.json
     */
    /**
     * Sample code: ApiManagementUserConfirmationPasswordSend.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUserConfirmationPasswordSend(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .userConfirmationPasswords()
            .sendWithResponse("rg1", "apimService1", "57127d485157a511ace86ae7", null, Context.NONE);
    }
}
