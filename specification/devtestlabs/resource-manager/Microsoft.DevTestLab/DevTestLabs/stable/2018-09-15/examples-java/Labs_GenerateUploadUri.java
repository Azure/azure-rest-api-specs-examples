
import com.azure.resourcemanager.devtestlabs.models.GenerateUploadUriParameter;

/**
 * Samples for Labs GenerateUploadUri.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Labs_GenerateUploadUri.json
     */
    /**
     * Sample code: Labs_GenerateUploadUri.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void labsGenerateUploadUri(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.labs().generateUploadUriWithResponse("resourceGroupName", "{labName}",
            new GenerateUploadUriParameter().withBlobName("{blob-name}"), com.azure.core.util.Context.NONE);
    }
}
