
import com.azure.resourcemanager.support.models.UploadFile;

/**
 * Samples for FilesNoSubscription Upload.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/UploadFile.json
     */
    /**
     * Sample code: UploadFile.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void uploadFile(com.azure.resourcemanager.support.SupportManager manager) {
        manager.filesNoSubscriptions().uploadWithResponse("testworkspaceName", "test.txt", new UploadFile().withContent(
            "iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAMAAAAoLQ9TAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAAgY0hSTQAAeiYAAICEAAD6AAAAgOgAAHUwAADqYAAAOpgAABd")
            .withChunkIndex(0), com.azure.core.util.Context.NONE);
    }
}
