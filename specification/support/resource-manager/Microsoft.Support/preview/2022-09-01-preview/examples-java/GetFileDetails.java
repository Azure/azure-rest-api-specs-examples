/** Samples for FilesNoSubscription Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/GetFileDetails.json
     */
    /**
     * Sample code: Get details of a subscription file.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void getDetailsOfASubscriptionFile(com.azure.resourcemanager.support.SupportManager manager) {
        manager.filesNoSubscriptions().getWithResponse("testworkspace", "test.txt", com.azure.core.util.Context.NONE);
    }
}
