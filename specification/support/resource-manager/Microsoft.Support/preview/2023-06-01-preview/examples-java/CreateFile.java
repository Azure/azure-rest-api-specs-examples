
import com.azure.resourcemanager.support.fluent.models.FileDetailsInner;

/**
 * Samples for FilesNoSubscription Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/CreateFile.json
     */
    /**
     * Sample code: Create a file workspace.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createAFileWorkspace(com.azure.resourcemanager.support.SupportManager manager) {
        manager.filesNoSubscriptions().createWithResponse("testworkspace", "test.txt",
            new FileDetailsInner().withChunkSize(41423).withFileSize(41423).withNumberOfChunks(1),
            com.azure.core.util.Context.NONE);
    }
}
