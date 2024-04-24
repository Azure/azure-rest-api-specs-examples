
/**
 * Samples for Files Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/GetFileDetailsForSubscription
     * .json
     */
    /**
     * Sample code: Get details of a subscription file.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void getDetailsOfASubscriptionFile(com.azure.resourcemanager.support.SupportManager manager) {
        manager.files().getWithResponse("testworkspace", "test.txt", com.azure.core.util.Context.NONE);
    }
}
