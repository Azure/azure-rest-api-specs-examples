
/**
 * Samples for FileWorkspacesNoSubscription Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * GetFileWorkspaceDetails.json
     */
    /**
     * Sample code: Get details of a file workspace.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void getDetailsOfAFileWorkspace(com.azure.resourcemanager.support.SupportManager manager) {
        manager.fileWorkspacesNoSubscriptions().getWithResponse("testworkspace", com.azure.core.util.Context.NONE);
    }
}
