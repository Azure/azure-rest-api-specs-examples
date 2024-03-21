
/**
 * Samples for FilesNoSubscription List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * ListFilesUnderFileWorkspace.json
     */
    /**
     * Sample code: List files under a workspace.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void listFilesUnderAWorkspace(com.azure.resourcemanager.support.SupportManager manager) {
        manager.filesNoSubscriptions().list("testworkspace", com.azure.core.util.Context.NONE);
    }
}
