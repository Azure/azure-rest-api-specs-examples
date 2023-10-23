/** Samples for Files List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListFilesForSubscriptionUnderFileWorkspace.json
     */
    /**
     * Sample code: List files under a workspace for a subscription.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void listFilesUnderAWorkspaceForASubscription(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.files().list("testworkspace", com.azure.core.util.Context.NONE);
    }
}
